package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/a-h/templ"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"github.com/shahinrahimi/booknest/pkg/auth"
	"github.com/shahinrahimi/booknest/pkg/book"
	"github.com/shahinrahimi/booknest/pkg/user"
	"github.com/shahinrahimi/booknest/store"
	"github.com/shahinrahimi/booknest/views/home"
)

func main() {

	logger := log.New(os.Stdout, "[BOOKNEST] ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		logger.Panic("Unable to locate .env file", err)
		os.Exit(1)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	secretKey := os.Getenv("SECRECT_KEY")
	if listenAddr == "" || secretKey == "" {
		logger.Panic("Unable to get environmental variables for server")
		os.Exit(1)
	}

	rootUsername := os.Getenv("ROOT_USERNAME")
	rootPassword := os.Getenv("ROOT_PASSWORD")
	if rootUsername == "" || rootPassword == "" {
		logger.Panic("Unable to get environmental variables for root user")
		os.Exit(1)
	}

	// create booknest store
	sqliteStore := store.NewSqliteStore(logger)
	// make sure to close db connection
	defer func() {
		if err := sqliteStore.Close(); err != nil {
			logger.Printf("Error closing database connection: %v", err)
		}
	}()

	// create tables
	sqliteStore.Init()
	// create root user if not exists
	sqliteStore.SetupRootAdmin(rootUsername, rootPassword)
	// create serve mux
	sm := mux.NewRouter()
	// create cookie store
	cs := sessions.NewCookieStore([]byte(secretKey))
	// create handlers
	// auth
	authH := auth.NewHandler(logger, sqliteStore, cs)
	// user
	userH := user.NewHandler(logger, sqliteStore)
	// book
	bookH := book.NewHandler(logger, sqliteStore)

	// register auth handlers to router
	postA := sm.Methods(http.MethodPost).Subrouter()
	postA.HandleFunc("/api/auth", authH.Login)
	postA.HandleFunc("/api/auth/logout", authH.Logout)

	// register user handlers to router
	getU := sm.Methods(http.MethodGet).Subrouter()
	getU.HandleFunc("/api/user", userH.ListAll)
	getU.HandleFunc("/api/user/{id}", userH.ListSingle)

	postU := sm.Methods(http.MethodPost).Subrouter()
	postU.HandleFunc("/api/user", userH.Create)
	postU.Use(userH.MiddlewareValidateUser)

	putU := sm.Methods(http.MethodPut).Subrouter()
	putU.HandleFunc("/api/user/{id}", userH.Update)
	putU.Use(userH.MiddlewareValidateUser)

	deleteU := sm.Methods(http.MethodDelete).Subrouter()
	deleteU.HandleFunc("/api/user/{id}", userH.Delete)

	// register book handlers to router
	getB := sm.Methods(http.MethodGet).Subrouter()
	getB.HandleFunc("/api/book", bookH.ListAll)
	getB.HandleFunc("/api/book/{id}", bookH.ListSingle)

	postB := sm.Methods(http.MethodPost).Subrouter()
	postB.HandleFunc("/api/book", bookH.Create)
	// postB.Use(authH.MiddlewareRequireAdmin)
	postB.Use(bookH.MiddlewareValidateBook)

	putB := sm.Methods(http.MethodPut).Subrouter()
	putB.HandleFunc("/api/book/{id}", bookH.Update)
	// putB.Use(authH.MiddlewareRequireAdmin)
	putU.Use(bookH.MiddlewareValidateBook)

	deleteB := sm.Methods(http.MethodDelete).Subrouter()
	deleteB.HandleFunc("/api/book/{id}", bookH.Delete)
	// deleteB.Use(authH.MiddlewareRequireAdmin)

	// handlers for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getD := sm.Methods(http.MethodGet).Subrouter()
	getD.Handle("/docs", sh)
	getD.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	sm.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	homeCom := home.Index()
	getV := sm.Methods(http.MethodGet).Subrouter()
	getV.Handle("/", templ.Handler(homeCom))

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
