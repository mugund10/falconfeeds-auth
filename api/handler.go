package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/mugund10/falconfeeds-auth/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// @Summary Health Check
// @Description Returns OK if the service is running
// @Tags Health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
// handler which simply writes ok to the response
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// @Summary Signup
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param signup body types.SignupRequest true "Signup Request"
// @Success 201 {object} types.User
// @Failure 400 {string} string "invalid json or validation error"
// @Failure 409 {string} string "user already exists"
// @Failure 500 {string} string "internal server error"
// @Router /signup [post]
// handler for signup
func (s *Server) handleSignup(w http.ResponseWriter, r *http.Request) {
	// context for db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
	if _, err := s.store.GetByEmail(ctx, sr.Email); err == nil {
		EncodeError(w, http.StatusConflict, "user already exists")
		return
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		EncodeError(w, http.StatusInternalServerError, "unable to connect to database")
		return
	}
	// 	storing user details to db
	user, err := types.NewUser(sr.Name, sr.Email, sr.Password)
	if err != nil {
		EncodeError(w, http.StatusInternalServerError, "user cant be created at this time")
		return
	}
	if err := s.store.Insert(ctx, user); err != nil {
		EncodeError(w, http.StatusInternalServerError, "user cant be created at this time")
		log.Println(err)
		return
	}
	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// @Summary Login
// @Description Login a user and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body types.LoginRequest true "Login Request"
// @Success 200 {object} types.LoginResponse
// @Failure 400 {string} string "invalid json or validation error"
// @Failure 401 {string} string "invalid email or password"
// @Failure 500 {string} string "internal server error"
// @Router /login [post]
// handler for login
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// context for db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// validating json fields
	var lr types.LoginRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&lr); err != nil {
		EncodeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	// validating json contents
	if err := lr.Validate(); err != nil {
		EncodeError(w, http.StatusBadRequest, err.Error())
		return
	}
	// finds user by email
	user, err := s.store.GetByEmail(ctx, lr.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			EncodeError(w, http.StatusUnauthorized, "invalid email or password")
		} else {
			EncodeError(w, http.StatusInternalServerError, "unable to connect to database")
		}
		return
	}
	// checks password
	if !user.ValidatePass(lr.Password) {
		EncodeError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}
	token, err := NewJwt(user.ID.Hex(), user.Email).Sign()
	if err != nil {
		EncodeError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}
	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types.LoginResponse{
		Token: token,
	})
}

// @Summary Test Endpoint
// @Description Protected route to verify JWT token
// @Tags Test
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string "missing or invalid token"
// @Router /test [get]
// handler for test
func (s *Server) handleTest(w http.ResponseWriter, r *http.Request) {
	// parsing token from header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		EncodeError(w, http.StatusUnauthorized, "missing token")
		return
	}
	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		EncodeError(w, http.StatusUnauthorized, "invalid token format")
		return
	}
	tokenStr := authHeader[len(prefix):]
	// validating token
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		EncodeError(w, http.StatusUnauthorized, err.Error())
		return
	}
	// response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "You are allowed",
		"userID":  claims.ID,
		"email":   claims.Email,
	})
}
