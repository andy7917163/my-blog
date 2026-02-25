## ADDED Requirements

### Requirement: Front Matter 解析
系統 SHALL 解析每篇文章頂部的 YAML Front Matter，支援欄位：`title`（字串）、`date`（日期）、`style`（字串，可選）。

#### Scenario: 標準 frontmatter 解析
- **WHEN** 文章包含合法 YAML frontmatter
- **THEN** 系統正確提取 title、date、style 欄位供模板使用

#### Scenario: 缺少 title
- **WHEN** frontmatter 缺少 `title` 欄位
- **THEN** 系統以檔名作為標題，並顯示警告

### Requirement: Markdown 轉 HTML
系統 SHALL 使用 goldmark 將 Markdown 內容轉換為 HTML，支援 CommonMark 規範。

#### Scenario: 標準 Markdown 轉換
- **WHEN** 文章包含標題、段落、清單、程式碼區塊
- **THEN** 產生對應的語意化 HTML

#### Scenario: 空文章
- **WHEN** Markdown 內容為空（只有 frontmatter）
- **THEN** 產生內容為空的文章頁面，不報錯

### Requirement: 文章排序
系統 SHALL 在文章列表頁依 `date` 欄位由新到舊排序文章。

#### Scenario: 依日期排序
- **WHEN** 多篇文章有不同 date
- **THEN** 列表頁依日期降冪排列
