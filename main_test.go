package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/booknest/sdk/client"
	"github.com/shahinrahimi/booknest/sdk/client/book"
)

func TestClient(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal("the environmental varaible not set correctly", err)
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		t.Fatal("the environmental varaible not set correctly")
	}
	cfg := client.DefaultTransportConfig().WithHost("localhost" + listenAddr)
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := book.NewListBooksParams()
	books, err := c.Book.ListBooks(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(books)

}
