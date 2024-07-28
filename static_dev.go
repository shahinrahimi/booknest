//+build dev
//go:build dev
// +build dev

package main

import (
	"net/http"
)

func staticFileHandler() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
}
