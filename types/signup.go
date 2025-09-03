package types

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
)

// Request body for signup
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response Body for signup
type SignupResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// validates the user input
func (srq *SignupRequest) Validate() error {
	// checks whether the fields contins empty spaces
	if srq.Name != strings.TrimSpace(srq.Name) ||
		srq.Email != strings.TrimSpace(srq.Email) ||
		srq.Password != strings.TrimSpace(srq.Password) {
		return fmt.Errorf("fields contains leading or trailing spaces")
	}
	// checks whether its in the mail format
	_, err := mail.ParseAddress(srq.Email)
	if err != nil {
		return fmt.Errorf("not a valid email address %s", err)
	}
	// checks whether the username or password is blank
	if len(srq.Name) < 4 || len(srq.Password) < 8 {
		return fmt.Errorf("username mustbe 4 charaters long and password must be 8 characters")
	}
	return nil
}
