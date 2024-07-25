package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	logger := log.New(os.Stdout, "[BOOKNEST] ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		logger.Panic("Unable to locate .env file", err)
		os.Exit(1)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		logger.Panic("Unable to get environmental variable")
		os.Exit(1)
	}

	sm := mux.NewRouter()

	s := http.Server{
		Addr:     listenAddr,
		Handler:  sm,
		ErrorLog: logger,
	}

	go func() {
		logger.Println("Starting server on port", listenAddr)
		if err := s.ListenAndServe(); err != nil {
			logger.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	defer cancel()
}
