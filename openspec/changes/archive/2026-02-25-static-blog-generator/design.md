## Context

從零打造靜態部落格產生器，無既有程式碼。技術棧：Go（引擎）+ 純 CSS/Vanilla JS（前端）。部署目標為 Vercel 靜態托管。

## Goals / Non-Goals

**Goals:**
- 單一 Go binary，包含 CLI 所有功能
- Markdown + Front Matter → HTML 的完整 pipeline
- 每篇文章可選擇性掛載自訂 CSS（via frontmatter `style` 欄位）
- 本地 `serve` 指令提供開發預覽
- 產出純靜態檔案，直接部署 Vercel

**Non-Goals:**
- 插件系統（第二階段）
- 多主題切換（第二階段）
- 標籤、分類、分頁、RSS（第二階段）
- 前端打包工具（Vite、Webpack 等）
- 資料庫或任何後端服務

## Decisions

### 1. Go 單一 binary 架構

所有功能（CLI、Markdown 解析、模板渲染、HTTP server）打包成單一 `blog` binary。

**理由**：部署與使用簡單，無需 Node.js runtime。Go 標準函式庫提供 HTTP server、file I/O、template engine，依賴極少。

### 2. Markdown 解析使用 goldmark

**理由**：goldmark 是 Go 生態中最活躍、符合 CommonMark 規範的 Markdown 解析器，可擴充（syntax highlighting 等未來需求）。替代方案 `blackfriday` 已不活躍維護。

### 3. 模板引擎使用 Go 內建 `html/template`

**理由**：零額外依賴，自動 XSS 防護，足以滿足靜態頁面需求。替代方案 `Pongo2`（Jinja2-like）彈性更高但增加複雜度，MVP 不需要。

### 4. 自訂 CSS 透過 frontmatter `style` 欄位注入

```yaml
---
title: 我的文章
date: 2026-02-25
style: styles/cyberpunk.css  # 相對於 posts/ 目錄
---
```

Build 時，若文章有 `style` 欄位，在輸出的 HTML `<head>` 中插入對應的 `<link rel="stylesheet">` 標籤。

**理由**：最簡單的實作方式，不需要特殊語法，純 frontmatter 宣告。

### 5. 輸出目錄結構

```
public/
├── index.html              # 文章列表
├── posts/
│   └── <slug>/
│       └── index.html      # 每篇文章（slug 由檔名產生）
├── styles/                 # 複製自 posts/styles/
└── themes/
    └── default/
        └── style.css       # 預設主題 CSS
```

每篇文章使用 `<slug>/index.html` 結構，讓 URL 不帶 `.html`（Vercel 自動處理）。

### 6. 專案目錄結構（Go module）

```
my-blog/
├── cmd/
│   └── blog/
│       └── main.go         # CLI 入口
├── internal/
│   ├── builder/            # build 邏輯
│   ├── parser/             # Markdown + frontmatter 解析
│   ├── server/             # 本地 HTTP server
│   └── theme/              # 模板管理
├── blog.yaml               # 全站設定
├── posts/                  # 使用者文章
│   └── styles/             # 文章自訂 CSS
├── themes/
│   └── default/
│       ├── list.html       # 文章列表模板
│       ├── post.html       # 文章內頁模板
│       └── style.css
├── public/                 # build 輸出（gitignore）
├── vercel.json             # Vercel 部署設定
└── go.mod
```

## Risks / Trade-offs

- **Go template 彈性有限** → 對 MVP 已足夠；若未來需要複雜邏輯可換 Pongo2
- **無 hot reload**（`serve` 只是靜態伺服器）→ 使用者需手動重跑 `build`；第二階段可加 file watcher
- **slug 衝突**（兩篇文章檔名相同）→ build 時報錯提示使用者

## Open Questions

- `blog.yaml` 要支援哪些全站設定欄位？（建議：`title`、`description`、`baseURL`）
- `serve` 是否需要自動監聽檔案變化並重新 build？（暫定：不需要，MVP 手動）
