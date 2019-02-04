package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(1, 3)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("{\"error\": \"Rate Limit Exceeded\"}"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
