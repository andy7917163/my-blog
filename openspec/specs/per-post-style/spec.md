## ADDED Requirements

### Requirement: 自訂 CSS 注入
系統 SHALL 在文章 frontmatter 包含 `style` 欄位時，於產生的 HTML `<head>` 中插入對應的 `<link rel="stylesheet">` 標籤，路徑相對於 `public/` 根目錄。

#### Scenario: 文章有自訂 style
- **WHEN** 文章 frontmatter 含 `style: styles/cyberpunk.css`
- **THEN** 產生的 HTML head 包含 `<link rel="stylesheet" href="/styles/cyberpunk.css">`

#### Scenario: 文章無自訂 style
- **WHEN** 文章 frontmatter 不含 `style` 欄位
- **THEN** HTML head 只有預設主題 CSS，不插入額外 stylesheet

### Requirement: 自訂 CSS 檔案複製
系統 SHALL 在 build 時，將 `posts/styles/` 目錄下的所有 CSS 檔案複製至 `public/styles/`。

#### Scenario: 複製自訂樣式
- **WHEN** `posts/styles/` 存在 CSS 檔案
- **THEN** build 後 `public/styles/` 包含相同檔案

#### Scenario: styles 目錄不存在
- **WHEN** `posts/styles/` 目錄不存在
- **THEN** build 正常完成，不報錯
