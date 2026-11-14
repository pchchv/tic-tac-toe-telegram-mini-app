package server

import (
	"net/http"
	"os"
	"path/filepath"
)

// spaHandler serves static assets from staticDir and falls back to index.html for unknown routes
func spaHandler(staticDir string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(staticDir))
	return func(w http.ResponseWriter, r *http.Request) {
		// Clean the requested path to prevent directory traversal
		cleanPath := filepath.Clean(r.URL.Path)
		fullPath := filepath.Join(staticDir, cleanPath)
		// Check if the requested file exists and is not a directory
		info, err := os.Stat(fullPath)
		if os.IsNotExist(err) || (err == nil && info.IsDir() && cleanPath != "/") {
			// Serve index.html for client-side routing
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}

		// Serve the existing static asset (JS, CSS, images, or root index.html)
		fileServer.ServeHTTP(w, r)
	}
}