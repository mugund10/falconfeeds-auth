package api

import "net/http"

type Server struct {
	listenAddr string
}

// Newserver creates a server with the given address
func Newserver(listenAddress string) *Server {
	return &Server{
		listenAddr: listenAddress,
	}
}

// starts and registers servers routes
func (s *Server) Start() error {
	http.HandleFunc("/health", s.handleHealth)
	return http.ListenAndServe(s.listenAddr, nil)
}

// handler which simply writes ok to the response
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
