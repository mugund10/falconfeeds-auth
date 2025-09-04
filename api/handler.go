package api

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"github.com/mugund10/falconfeeds-auth/types"
)

// handler which simply writes ok to the response
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// handler for signup
func (s *Server) handleSignup(w http.ResponseWriter, r *http.Request) {
	// validating json fields
	var sr types.SignupRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&sr); err != nil {
		EncodeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	// validating json contents
	if err := sr.Validate(); err != nil {
		EncodeError(w, http.StatusBadRequest, err.Error())
		return
	}
	// checking user with same email id
	// if _, err := s.store.GetByEmail(context.TODO(), sr.Email); err == nil {
	// 	EncodeError(w, http.StatusConflict, "user already exists")
	// 	return
	// }
	// todo : create a method for or generating a user
	// var user types.User
	//storing user details to db
	// if err := s.store.Insert(context.TODO(), user); err != nil {
	// 	EncodeError(w, http.StatusInternalServerError, "user cant be created at this time")
	// }
	tt := time.Now().UTC()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(types.SignupResponse{Id: "10", Name: sr.Name, Email: sr.Email, CreatedAt: tt})

	fmt.Printf("%+v\n", sr)
}

// handler for login
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// encodes error to responsewriter
func EncodeError(w http.ResponseWriter, statusCode int, Content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.ErrorResponse{Error: Content})
}
