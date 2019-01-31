package session

import (
	"./structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	boardId, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	session := structs.OutDTO{UUID:id, Board:boardId}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Println(err)
		return
	}

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
