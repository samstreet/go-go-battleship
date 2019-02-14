package core

import (
	UserServices "../../services"
	"net/http"
)

func HeaderValidUserToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		identifier := r.Header.Get("X-USER-IDENTIFIER")
		userService := UserServices.NewUserService()
		_, err := userService.FindByUUID(identifier)

		if err != nil {
			w.Header().Set("Content-Type", r.Header.Get("Accept"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
