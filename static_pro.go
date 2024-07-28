//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public/*
var publicFiles embed.FS

func staticFileHandler() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.FS(publicFiles)))
}
