package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Serve 在指定 port 提供靜態檔案服務
func Serve(publicDir string, port int) error {
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		return fmt.Errorf("找不到 %s 目錄，請先執行 blog build", publicDir)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(publicDir, filepath.Clean(r.URL.Path))

		// 若為目錄，嘗試 index.html
		info, err := os.Stat(path)
		if err == nil && info.IsDir() {
			indexPath := filepath.Join(path, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
		}

		// 防止路徑逸出 publicDir
		if !strings.HasPrefix(path, filepath.Clean(publicDir)) {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}

		http.ServeFile(w, r, path)
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("🚀 本地預覽：http://localhost%s\n", addr)
	return http.ListenAndServe(addr, mux)
}
