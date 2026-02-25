package builder

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/andyhuang/my-blog/internal/parser"
	"github.com/andyhuang/my-blog/internal/theme"
)

// Config 代表 build 所需的設定
type Config struct {
	PostsDir  string
	ThemeDir  string
	OutputDir string
	Site      theme.SiteConfig
}

// Build 執行完整的靜態網站建構流程
func Build(cfg Config) error {
	// 1. 掃描所有 .md 文章
	posts, err := scanPosts(cfg.PostsDir)
	if err != nil {
		return fmt.Errorf("掃描文章失敗: %w", err)
	}

	// 2. 檢查 slug 衝突
	if err := checkSlugConflicts(posts); err != nil {
		return err
	}

	// 3. 依日期排序
	parser.SortByDateDesc(posts)

	// 4. 載入主題
	t, err := theme.Load(cfg.ThemeDir)
	if err != nil {
		return fmt.Errorf("載入主題失敗: %w", err)
	}

	// 5. 清空並重建輸出目錄
	if err := os.RemoveAll(cfg.OutputDir); err != nil {
		return fmt.Errorf("清除輸出目錄失敗: %w", err)
	}

	// 6. 產生文章頁
	for _, post := range posts {
		if err := renderPost(cfg, t, post); err != nil {
			return fmt.Errorf("渲染文章 %s 失敗: %w", post.Slug, err)
		}
	}

	// 7. 產生列表頁
	if err := renderList(cfg, t, posts); err != nil {
		return fmt.Errorf("渲染列表頁失敗: %w", err)
	}

	// 8. 複製自訂樣式
	if err := copyStyles(cfg.PostsDir, cfg.OutputDir); err != nil {
		return fmt.Errorf("複製樣式失敗: %w", err)
	}

	// 9. 複製主題 CSS
	if err := copyThemeAssets(cfg.ThemeDir, cfg.OutputDir); err != nil {
		return fmt.Errorf("複製主題資源失敗: %w", err)
	}

	fmt.Printf("✓ Build 完成，共 %d 篇文章，輸出至 %s\n", len(posts), cfg.OutputDir)
	return nil
}

func scanPosts(postsDir string) ([]*parser.Post, error) {
	entries, err := filepath.Glob(filepath.Join(postsDir, "*.md"))
	if err != nil {
		return nil, err
	}
	var posts []*parser.Post
	for _, path := range entries {
		post, err := parser.ParseFile(path)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func checkSlugConflicts(posts []*parser.Post) error {
	seen := make(map[string]string)
	for _, p := range posts {
		if prev, ok := seen[p.Slug]; ok {
			return fmt.Errorf("slug 衝突：%s 和 %s 產生相同的 slug %q，請重新命名", prev, p.Source, p.Slug)
		}
		seen[p.Slug] = p.Source
	}
	return nil
}

func renderPost(cfg Config, t *theme.Theme, post *parser.Post) error {
	dir := filepath.Join(cfg.OutputDir, "posts", post.Slug)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.RenderPost(&buf, cfg.Site, post); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "index.html"), buf.Bytes(), 0644)
}

func renderList(cfg Config, t *theme.Theme, posts []*parser.Post) error {
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.RenderList(&buf, cfg.Site, posts); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(cfg.OutputDir, "index.html"), buf.Bytes(), 0644)
}

func copyStyles(postsDir, outputDir string) error {
	src := filepath.Join(postsDir, "styles")
	dst := filepath.Join(outputDir, "styles")

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil // styles 目錄不存在，跳過
	}

	return copyDir(src, dst)
}

func copyThemeAssets(themeDir, outputDir string) error {
	themeName := filepath.Base(themeDir)
	dst := filepath.Join(outputDir, "themes", themeName)
	return copyDir(themeDir, dst)
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
