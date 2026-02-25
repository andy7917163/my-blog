package theme

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/andyhuang/my-blog/internal/parser"
)

// SiteConfig 代表全站設定
type SiteConfig struct {
	Title       string
	Description string
	BaseURL     string
}

// Theme 代表一個主題，持有已載入的模板
type Theme struct {
	listTmpl *template.Template
	postTmpl *template.Template
}

// Load 從指定目錄載入主題模板
func Load(themeDir string) (*Theme, error) {
	listPath := filepath.Join(themeDir, "list.html")
	postPath := filepath.Join(themeDir, "post.html")

	// html/template 自動 HTML escape，安全
	listTmpl, err := template.New("list").Funcs(templateFuncs()).ParseFiles(listPath)
	if err != nil {
		return nil, err
	}
	postTmpl, err := template.New("post").Funcs(templateFuncs()).ParseFiles(postPath)
	if err != nil {
		return nil, err
	}

	return &Theme{listTmpl: listTmpl, postTmpl: postTmpl}, nil
}

// RenderList 渲染文章列表頁
func (t *Theme) RenderList(w io.Writer, site SiteConfig, posts []*parser.Post) error {
	data := map[string]any{
		"Site":  site,
		"Posts": posts,
	}
	return t.listTmpl.ExecuteTemplate(w, "list.html", data)
}

// RenderPost 渲染單篇文章頁
func (t *Theme) RenderPost(w io.Writer, site SiteConfig, post *parser.Post) error {
	data := map[string]any{
		"Site": site,
		"Post": struct {
			*parser.Post
			Content template.HTML
		}{
			Post:    post,
			Content: template.HTML(post.Content),
		},
	}
	return t.postTmpl.ExecuteTemplate(w, "post.html", data)
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{}
}
