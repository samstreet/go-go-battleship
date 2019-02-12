package middleware

import (
	SessionServices "../../../session/services"
	UserServices "../../services"
	"github.com/gorilla/mux"
	"net/http"
)

func UserIsAttachedToBoard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		identifier := r.Header.Get("X-USER-IDENTIFIER")
		vars := mux.Vars(r)
		board := vars["uuid"]

		userService := UserServices.NewUserService()

		user, _ := userService.FindByUUID(identifier)
		session := SessionServices.NewSessionService().FindSessionByUUID(user.SessionID)

		if session.Board.UUID.String() != board {
			w.Header().Set("Content-Type", r.Header.Get("Accept"))
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
