package api

import (
	"encoding/json"
	"net/http"

	"github.com/mugund10/falconfeeds-auth/types"
)

// encodes error to responsewriter
func EncodeError(w http.ResponseWriter, statusCode int, Content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.ErrorResponse{Error: Content})
}
