package core

import (
	BoardModels "../../../board/model"
	BoardServices "../../../board/services"
	"github.com/gorilla/mux"
	"net/http"
)

func BoardExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boardService := BoardServices.NewBoardService(BoardModels.BoardModel{})

		_, error := boardService.FindByUUID(vars["uuid"])

		if error != nil {
			w.Header().Set("Content-Type", r.Header.Get("Accept"))
			body := ErrorResponse{Message: error.Error()}
			w.Write(ErrorToSchema(body, r.Header.Get("Accept")))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
