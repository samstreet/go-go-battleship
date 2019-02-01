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
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	helpers.HandleError(err)

	boardId, err := uuid.NewV4()
	helpers.HandleError(err)

	board := BoardStructs.BoardOutDTO{UUID:boardId}
	session := structs.SessionOutDTO{UUID:id, Board: board}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(session)
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
