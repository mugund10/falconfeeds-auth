package api

import (
	"net/http"
	"time"
)

// a simple rate limiter
func SimpleRateLimiter(maxRequests int, per time.Duration) func(http.Handler) http.Handler {
	throttle := time.Tick(per / time.Duration(maxRequests))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			<-throttle
			next.ServeHTTP(w, r)
		})
	}
}
