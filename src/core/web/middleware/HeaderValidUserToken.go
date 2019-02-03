package middleware

import (
	"log"
	"net/http"
)

func HeaderValidUserToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Check Token is a valid token")
		h.ServeHTTP(w, r)
	})
}