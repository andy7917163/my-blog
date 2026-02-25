## ADDED Requirements

### Requirement: 靜態檔案服務
系統 SHALL 提供本地 HTTP server，服務 `public/` 目錄下的靜態檔案，預設埠號 `3000`。

#### Scenario: 正常啟動
- **WHEN** 使用者執行 `blog serve` 且 `public/` 存在
- **THEN** server 在 `http://localhost:3000` 啟動，終端顯示網址

#### Scenario: public 目錄不存在
- **WHEN** `public/` 目錄不存在
- **THEN** 系統顯示錯誤提示，建議先執行 `blog build`

### Requirement: 目錄 index 處理
系統 SHALL 自動將目錄請求導向對應的 `index.html`，使 `/posts/hello/` 正確返回 `posts/hello/index.html`。

#### Scenario: 目錄請求
- **WHEN** 瀏覽器請求 `/posts/hello/`
- **THEN** server 返回 `public/posts/hello/index.html` 內容
