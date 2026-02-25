## ADDED Requirements

### Requirement: 預設主題模板
系統 SHALL 提供預設主題，包含文章列表頁模板（`list.html`）與文章內頁模板（`post.html`），使用 Go `html/template` 語法。

#### Scenario: 渲染文章列表
- **WHEN** 執行 build
- **THEN** `public/index.html` 使用 `list.html` 模板，列出所有文章標題、日期與連結

#### Scenario: 渲染文章內頁
- **WHEN** 執行 build
- **THEN** `public/posts/<slug>/index.html` 使用 `post.html` 模板，包含文章標題、日期、HTML 內容

### Requirement: 預設主題 CSS
系統 SHALL 提供基礎的預設 CSS（`themes/default/style.css`），套用至所有頁面，確保基本可讀性。

#### Scenario: 預設樣式套用
- **WHEN** 文章未指定自訂 style
- **THEN** 頁面套用預設主題 CSS，呈現合理的排版與字型

### Requirement: 全站設定傳入模板
系統 SHALL 將 `blog.yaml` 的設定（`title`、`description`、`baseURL`）傳入所有模板，供頁面標題等使用。

#### Scenario: 全站標題顯示
- **WHEN** `blog.yaml` 設定 `title: 我的部落格`
- **THEN** 所有頁面 `<title>` 包含此站名
