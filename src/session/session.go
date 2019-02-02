package session

import (
	BoardStructs "../board/structs"
	"../core/helpers"
	"./structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"../session/services"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionService := services.NewSessionService()

	session := sessionService.CreateSession()
	board := BoardStructs.BoardOutDTO{UUID:uuid.FromStringOrNil(session.Board.ID)}

	sessionOut := structs.SessionOutDTO{UUID:uuid.FromStringOrNil(session.ID), Board: board}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(sessionOut)
	helpers.HandleError(err)

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"uuid\": \"" + id.String() + "\"}"))
	w.WriteHeader(http.StatusOK)
}

func ViewSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"uuid\": \"" + vars["session"] + "\"}"))
	w.WriteHeader(http.StatusOK)
}
