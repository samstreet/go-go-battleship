package middleware

import (
	"net/http"
)

func HeaderUserTokenValidator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identifier := r.Header.Get("X-USER-IDENTIFIER")
		if identifier == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"error\": \"Missing X-USER-IDENTIFIER Header\"}"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
