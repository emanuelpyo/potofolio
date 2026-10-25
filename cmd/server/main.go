package main

import (
	"log"
	"net/http"
	"time"

	"deryl-portfolio/internal/chat"
	"deryl-portfolio/internal/config"
	"deryl-portfolio/internal/site"
)

func main() {
	cfg := config.Load()

	page, err := site.New(cfg.BaseURL)
	if err != nil {
		log.Fatalf("render site: %v", err)
	}

	static := http.FileServer(http.Dir(cfg.StaticDir))

	mux := http.NewServeMux()
	mux.Handle("/api/chat", chat.New(cfg))
	for _, prefix := range []string{"/css/", "/js/", "/fonts/", "/images/", "/icons/", "/projects/", "/cv/"} {
		mux.Handle(prefix, static)
	}
	mux.Handle("/", page)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("Deryl Portfolio (Go) listening on http://localhost:%s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server: %v", err)
	}
}
