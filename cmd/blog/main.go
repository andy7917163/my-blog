package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/andyhuang/my-blog/internal/builder"
	"github.com/andyhuang/my-blog/internal/server"
	"github.com/andyhuang/my-blog/internal/theme"
	"gopkg.in/yaml.v3"
)

// SiteConfig 對應 blog.yaml
type SiteConfig struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	BaseURL     string `yaml:"baseURL"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		runBuild()
	case "serve":
		runServe()
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("用法：blog new <標題>")
			os.Exit(1)
		}
		runNew(strings.Join(os.Args[2:], " "))
	default:
		fmt.Printf("未知指令：%s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`用法：blog <指令>

指令：
  new <標題>   建立新文章
  build        產生靜態網站
  serve        啟動本地預覽伺服器`)
}

func loadSiteConfig() (SiteConfig, error) {
	data, err := os.ReadFile("blog.yaml")
	if err != nil {
		return SiteConfig{Title: "My Blog"}, nil
	}
	var cfg SiteConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return SiteConfig{}, fmt.Errorf("解析 blog.yaml 失敗: %w", err)
	}
	return cfg, nil
}

func runBuild() {
	site, err := loadSiteConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "錯誤:", err)
		os.Exit(1)
	}

	cfg := builder.Config{
		PostsDir:  "posts",
		ThemeDir:  filepath.Join("themes", "default"),
		OutputDir: "public",
		Site: theme.SiteConfig{
			Title:       site.Title,
			Description: site.Description,
			BaseURL:     site.BaseURL,
		},
	}

	if err := builder.Build(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "Build 失敗:", err)
		os.Exit(1)
	}
}

func runServe() {
	if err := server.Serve("public", 3000); err != nil {
		fmt.Fprintln(os.Stderr, "錯誤:", err)
		os.Exit(1)
	}
}

func runNew(title string) {
	slug := titleToSlug(title)
	filename := fmt.Sprintf("%s-%s.md", time.Now().Format("2006-01-02"), slug)
	path := filepath.Join("posts", filename)

	if _, err := os.Stat(path); err == nil {
		fmt.Fprintf(os.Stderr, "檔案已存在：%s\n", path)
		os.Exit(1)
	}

	content := fmt.Sprintf("---\ntitle: %s\ndate: %s\n---\n\n在這裡開始寫作...\n",
		title, time.Now().Format("2006-01-02"))

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "建立文章失敗:", err)
		os.Exit(1)
	}

	fmt.Printf("✓ 已建立文章：%s\n", path)
}

func titleToSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return '-'
	}, slug)
	// 壓縮連續的 -
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	return strings.Trim(slug, "-")
}
