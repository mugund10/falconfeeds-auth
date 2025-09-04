package api

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(GetEnv("JWT_SECRET", "Mugund10TEST").(string))

// a custom claim
type jWT struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewJwt(id, email string) *jWT {
	return &jWT{ID: id, Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

// creates token with claims
func (j *jWT) Sign() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	return token.SignedString(jwtKey)
}

// validates jwt
func ValidateToken(tokenStr string) (*jWT, error) {
	claims := &jWT{}
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	}
	// Parsing token
	token, err := jwt.ParseWithClaims(tokenStr, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	// verifies token
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("expired token")
	}
	return claims, nil
}
