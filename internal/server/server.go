// Package server handles the HTTP server setup and initialization
package server

import (
	"ascii-art-web-stylize/internal/config"
	"ascii-art-web-stylize/internal/handlers"
	"net/http"
)

// Server wraps the standard http.Server
type Server struct {
	server *http.Server
}

// Create sets up a new server instance with routes for static files, home page and ASCII art generation
func Create(cfg *config.Config) *Server {
	mux := http.NewServeMux()
	h := handlers.New(cfg)
	fs := http.FileServer(http.Dir(cfg.StaticPath))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/generate", h.Generate)
	mux.HandleFunc("/export", h.Export)
	return &Server{
		server: &http.Server{
			Addr:    cfg.Port,
			Handler: mux,
		},
	}
}

// Start begins listening for HTTP requests
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
