package api

import (
	"net/http"
	"time"

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
	limiter := SimpleRateLimiter(GetEnv("RATE_PER_MIN", 60).(int), time.Minute)
	// custom middleware stack
	mstack := MakeStack(limiter)
	// custom multiplexer
	mux := http.NewServeMux()
	// handler func
	mux.HandleFunc("GET  /health", s.handleHealth)
	mux.HandleFunc("POST /signup", s.handleSignup)
	mux.HandleFunc("POST /login", s.handleLogin)
	mux.HandleFunc("GET  /test", s.handleTest)
	// custom server
	server := http.Server{
		Addr:    s.listenAddr,
		Handler: mstack(mux),
	}
	return server.ListenAndServe()
}
