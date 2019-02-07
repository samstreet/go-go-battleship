package middleware

import (
	"net/http"
)

func HeaderValidUserToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)
		return

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"Invalid X-USER-IDENTIFIER\"}"))
		return
	})
}
