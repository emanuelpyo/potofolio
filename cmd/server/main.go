package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"deryl-portfolio/internal/chat"
	"deryl-portfolio/internal/config"
	"deryl-portfolio/internal/site"
)

func main() {
	cfg := config.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}
	if port == "" {
		port = "3000"
	}

	page, err := site.New(cfg.BaseURL)
	if err != nil {
		log.Fatalf("render site: %v", err)
	}

	static := http.FileServer(http.Dir(cfg.StaticDir))

	mux := http.NewServeMux()
	mux.Handle("/api/chat", chat.New(cfg))

	for _, prefix := range []string{
		"/css/",
		"/js/",
		"/fonts/",
		"/images/",
		"/icons/",
		"/projects/",
		"/cv/",
	} {
		mux.Handle(prefix, static)
	}

	mux.Handle("/", page)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("Deryl Portfolio (Go) listening on port %s", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server: %v", err)
	}
}
