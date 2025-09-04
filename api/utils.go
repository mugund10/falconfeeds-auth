package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mugund10/falconfeeds-auth/types"
)

// encodes error to responsewriter
func EncodeError(w http.ResponseWriter, statusCode int, Content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.ErrorResponse{Error: Content})
}

// checks for env values if not present returns fallback
func GetEnv(key string, fallback any) any {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
