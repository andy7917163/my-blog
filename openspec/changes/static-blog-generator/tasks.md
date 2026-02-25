## 1. 專案初始化

- [ ] 1.1 初始化 Go module（`go mod init`），建立基本目錄結構（`cmd/blog/`、`internal/`）
- [ ] 1.2 加入 `goldmark` 依賴（`go get github.com/yuin/goldmark`）
- [ ] 1.3 建立 `blog.yaml` 範例設定檔（含 `title`、`description`、`baseURL`）
- [ ] 1.4 初始化 git repo，設定 `.gitignore`（忽略 `public/`）

## 2. Markdown + Front Matter 解析（content-pipeline）

- [ ] 2.1 實作 `internal/parser`：使用 `gopkg.in/yaml.v3` 解析 YAML frontmatter
- [ ] 2.2 實作 goldmark Markdown → HTML 轉換
- [ ] 2.3 實作文章 slug 從檔名產生邏輯
- [ ] 2.4 實作文章依 `date` 降冪排序

## 3. 主題模板系統（theme-system）

- [ ] 3.1 建立 `themes/default/list.html`：文章列表頁模板
- [ ] 3.2 建立 `themes/default/post.html`：文章內頁模板（含自訂 style 注入點）
- [ ] 3.3 建立 `themes/default/style.css`：基礎預設樣式
- [ ] 3.4 實作 `internal/theme`：載入並渲染 Go html/template 模板

## 4. Build Pipeline（content-pipeline + per-post-style）

- [ ] 4.1 實作 `internal/builder`：掃描 `posts/` 所有 `.md` 檔案
- [ ] 4.2 實作 slug 衝突檢查，衝突時報錯終止
- [ ] 4.3 實作產生 `public/posts/<slug>/index.html`（套入 post.html 模板）
- [ ] 4.4 實作產生 `public/index.html`（套入 list.html 模板）
- [ ] 4.5 實作 `posts/styles/` → `public/styles/` 複製邏輯
- [ ] 4.6 實作 `themes/default/style.css` → `public/themes/default/style.css` 複製

## 5. CLI 指令（cli）

- [ ] 5.1 實作 `cmd/blog/main.go`：使用 `flag` 或 `os.Args` 解析子指令
- [ ] 5.2 實作 `blog build` 指令：呼叫 builder pipeline
- [ ] 5.3 實作 `blog new <title>` 指令：產生新文章 .md 檔含 frontmatter
- [ ] 5.4 實作 `blog serve` 指令：啟動本地 HTTP server（port 3000）

## 6. 本地 Server（local-server）

- [ ] 6.1 實作 `internal/server`：使用 `net/http` 服務 `public/` 靜態檔案
- [ ] 6.2 處理目錄請求自動導向 `index.html`
- [ ] 6.3 `public/` 不存在時顯示友善錯誤訊息

## 7. Vercel 部署設定

- [ ] 7.1 建立 `vercel.json`，設定靜態輸出目錄為 `public/`
- [ ] 7.2 驗證部署流程：本地 build → push → Vercel 自動部署
