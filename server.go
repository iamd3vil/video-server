package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
)

// Server is the main server to start
type Server struct {
	cfg    Config
	router *chi.Mux
}

// NewServer returns a server instance
func NewServer(cfg Config) (*Server, error) {
	router := chi.NewRouter()
	return &Server{
		cfg:    cfg,
		router: router,
	}, nil
}

// ListenAndServe starts the HTTP server
func (srv *Server) ListenAndServe(ctx context.Context) {
	srv.routes()
	s := &http.Server{
		Addr:    srv.cfg.HTTPAddr,
		Handler: srv.router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Println("exiting...")
			s.Shutdown(context.Background())
			wg.Done()
			break
		}
	}(ctx)

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				wg.Done()
				return
			}
			log.Fatalln(err)
		}
	}()

	log.Printf("Starting server on %s", srv.cfg.HTTPAddr)

	wg.Wait()
}

// Index serves just a welcome message
func (srv *Server) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	}
}
