package middleware

import (
	"net/http"
)

func HeaderUserTokenValidator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identifier := r.Header.Get("X-USER-TOKEN_IDENTIFIER")
		if identifier == "" {
			w.WriteHeader(http.StatusPreconditionRequired)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"error\": \"Missing X-USER-TOKEN_IDENTIFIER Header\"}"))
			return
		}

		h.ServeHTTP(w, r)
	})
}