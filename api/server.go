package api

import (
	"net/http"

	"github.com/mugund10/falconfeeds-auth/storage"
)

type Server struct {
	listenAddr string
	store      storage.UserStorer
}

// Newserver creates a server with the given address
func Newserver(listenAddress string, store storage.UserStorer) *Server {
	return &Server{
		listenAddr: listenAddress,
		store:      store,
	}
}

// starts and registers servers routes
func (s *Server) Start() error {
	http.HandleFunc("GET  /health", s.handleHealth)
	http.HandleFunc("POST /signup", s.handleSignup)
	http.HandleFunc("POST /login", s.handleLogin)
	return http.ListenAndServe(s.listenAddr, nil)
}
