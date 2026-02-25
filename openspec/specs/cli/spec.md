## ADDED Requirements

### Requirement: blog new 指令
系統 SHALL 提供 `blog new <title>` 指令，在 `posts/` 目錄下建立新的 Markdown 文章檔案，檔名由標題自動轉換為 kebab-case slug，並預填基本 Front Matter。

#### Scenario: 建立新文章
- **WHEN** 使用者執行 `blog new "我的第一篇文章"`
- **THEN** 系統在 `posts/` 建立 `my-di-yi-pian-wen-zhang.md`（或以日期前綴），內含 `title`、`date` frontmatter

#### Scenario: 標題含特殊字元
- **WHEN** 標題含空格或特殊字元
- **THEN** slug 自動轉為小寫 kebab-case，移除不合法字元

### Requirement: blog build 指令
系統 SHALL 提供 `blog build` 指令，讀取 `posts/` 所有 `.md` 檔案，產生靜態 HTML 至 `public/` 目錄。

#### Scenario: 成功 build
- **WHEN** 使用者執行 `blog build`
- **THEN** `public/` 包含 `index.html`（文章列表）與每篇文章的 `posts/<slug>/index.html`

#### Scenario: slug 衝突
- **WHEN** 兩篇文章產生相同 slug
- **THEN** 系統報錯並終止 build，提示使用者重新命名

### Requirement: blog serve 指令
系統 SHALL 提供 `blog serve` 指令，在本地啟動靜態 HTTP server 提供 `public/` 目錄內容預覽。

#### Scenario: 啟動本地伺服器
- **WHEN** 使用者執行 `blog serve`
- **THEN** 系統在 `http://localhost:3000` 提供靜態檔案服務，並在終端顯示網址
