package session

import (
	"../core/helpers"
	"../session/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionService := services.NewSessionService()
	session := sessionService.CreateSession()

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

	sessionService := services.NewSessionService()
	sessionOut := sessionService.FindSessionByUUID(vars["session"])

	b, err := json.Marshal(sessionOut)
	helpers.HandleError(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
