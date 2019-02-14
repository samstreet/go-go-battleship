package core

import (
	"net/http"
)

func HasUserTokenIdentifierHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-USER-IDENTIFIER") == "" {
			body := ErrorToSchema(ErrorResponse{Message: "Missing X-USER-IDENTIFIER Header"}, r.Header.Get("Accept"))
			w.Header().Set("Content-Type", r.Header.Get("Accept"))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			return
		}

		h.ServeHTTP(w, r)
	})
}
