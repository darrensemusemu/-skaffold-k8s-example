package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Hello at: " + t.String()))
	})

	addr := fmt.Sprintf(":%d", cfg.port)
	log.Printf("Running of port %s", addr)
	err := http.ListenAndServe(addr, r)
	return err
}
