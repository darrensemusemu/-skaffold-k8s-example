package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	port int
}

func main() {
	cfg := config{}
	flag.IntVar(&cfg.port, "port", 8080, "service port number, default 8080")
	flag.Parse()

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health/alive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	addr := fmt.Sprintf(":%d", cfg.port)
	log.Printf("Running of port %s", addr)
	http.ListenAndServe(addr, r)
	return nil
}
