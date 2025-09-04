package types

import (
	"fmt"
	"net/mail"
	"strings"
)

// Request body for Login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response body for Login
type LoginResponse struct {
	Token string `json:"token"`
}

// validates the user input
func (lrq *LoginRequest) Validate() error {
	// checks whether the fields contins empty spaces
	if lrq.Email != strings.TrimSpace(lrq.Email) ||
		lrq.Password != strings.TrimSpace(lrq.Password) {
		return fmt.Errorf("fields contains leading or trailing spaces")
	}
	// checks whether its in the mail format
	_, err := mail.ParseAddress(lrq.Email)
	if err != nil {
		return fmt.Errorf("not a valid email address %s", err)
	}
	// checks whether the password is in valid length
	if len(lrq.Password) < 8 {
		return fmt.Errorf("password must be 8 characters long")
	}
	return nil
}
