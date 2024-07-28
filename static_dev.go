//+build dev
//go:build dev
// +build dev

package main

import (
	"net/http"
	"path/filepath"
)

// mimeTypes maps file extensions to their MIME types.
var mimeTypes = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".html": "text/html",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".gif":  "image/gif",
	".svg":  "image/svg+xml",
	".ico":  "image/x-icon",
}

func fileServerWithMime(directory string) http.Handler {
	fs := http.FileServer(http.Dir(directory))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)
		w.Header().Set("Content-Type", "text/css")
		// if mimeType, ok := mimeTypes[ext]; ok {
		// 	w.Header().Set("Content-Type", mimeType)
		// }
		fs.ServeHTTP(w, r)
	})
}

func staticFileHandler() http.Handler {
	// return http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	return http.StripPrefix("/public/", fileServerWithMime("./public/"))
}
