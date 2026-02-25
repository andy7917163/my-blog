## Why

目前市面上的靜態部落格產生器（如 Hexo）功能龐大但彈性不足，無法讓每篇文章擁有完全獨立的視覺風格。本專案從零打造一個輕量、可控的靜態部落格引擎，以 Go 為核心，支援每篇文章透過 frontmatter 指定自訂 CSS，部署至 Vercel。

## What Changes

- 全新 CLI 工具 `blog`，支援 `new`、`build`、`serve` 指令
- Markdown 文章系統，支援 Front Matter（含 `style` 欄位）
- 每篇文章可選擇性掛載自訂 CSS 檔案
- Go `html/template` 模板引擎驅動的主題系統（預設主題）
- 純靜態輸出（HTML + CSS + Vanilla JS），可直接部署至 Vercel

## Capabilities

### New Capabilities

- `cli`: `blog new`、`blog build`、`blog serve` 三個核心 CLI 指令
- `content-pipeline`: Markdown 解析、Front Matter 處理、HTML 輸出
- `per-post-style`: 每篇文章透過 frontmatter `style` 欄位掛載自訂 CSS
- `theme-system`: 預設主題，包含文章列表頁與文章內頁的 HTML 模板與基礎 CSS
- `local-server`: 本地預覽伺服器，支援靜態檔案服務

### Modified Capabilities

## Impact

- 新增 Go module（`go.mod`），依賴 `goldmark`（Markdown 解析）
- 產出純靜態檔案至 `public/`，搭配 `vercel.json` 設定部署
- 無前端打包工具，CSS 與 JS 直接作為靜態資源輸出
