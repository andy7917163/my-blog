package parser

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

// FrontMatter 代表文章的 YAML 前置資料
type FrontMatter struct {
	Title string    `yaml:"title"`
	Date  time.Time `yaml:"date"`
	Style string    `yaml:"style"`
}

// Post 代表一篇解析完成的文章
type Post struct {
	FrontMatter
	Slug    string
	Content string // 已轉換的 HTML
	Source  string // 原始 Markdown 路徑
}

var frontMatterRegex = regexp.MustCompile(`(?s)^---\n(.*?)\n---\n?(.*)$`)

// ParseFile 解析單一 Markdown 檔案，回傳 Post
func ParseFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("讀取檔案失敗 %s: %w", path, err)
	}

	content := string(data)
	var fm FrontMatter
	var body string

	matches := frontMatterRegex.FindStringSubmatch(content)
	if matches != nil {
		if err := yaml.Unmarshal([]byte(matches[1]), &fm); err != nil {
			return nil, fmt.Errorf("解析 frontmatter 失敗 %s: %w", path, err)
		}
		body = strings.TrimSpace(matches[2])
	} else {
		body = content
	}

	// 若無 title，使用檔名
	if fm.Title == "" {
		base := filepath.Base(path)
		fm.Title = strings.TrimSuffix(base, filepath.Ext(base))
		fmt.Printf("警告：%s 缺少 title，使用檔名作為標題\n", path)
	}

	// Markdown → HTML
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(body), &buf); err != nil {
		return nil, fmt.Errorf("Markdown 轉換失敗 %s: %w", path, err)
	}

	slug := fileNameToSlug(filepath.Base(path))

	return &Post{
		FrontMatter: fm,
		Slug:        slug,
		Content:     buf.String(),
		Source:      path,
	}, nil
}

// fileNameToSlug 將檔名（不含副檔名）轉為 URL-safe slug
func fileNameToSlug(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	// 轉小寫，非 ASCII 字母數字和連字符替換為連字符
	var b strings.Builder
	prevDash := false
	for _, r := range strings.ToLower(name) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
		} else if !prevDash {
			b.WriteRune('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

// SortByDateDesc 依日期降冪排序文章
func SortByDateDesc(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
}
