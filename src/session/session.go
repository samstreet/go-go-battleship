package session

import (
	"../core/helpers"
	"../session/services"
	SessionDTO "../session/structs"
	"encoding/json"
	"encoding/xml"
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
	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.Write(transformDTOToSchema(sessionOut, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
	return
}

func transformDTOToSchema(dto SessionDTO.SessionOutDTO, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(dto)
		helpers.HandleError(err)
		return b
	}

	b, err := xml.MarshalIndent(dto, "  ", "    ")
	log.Print(b)
	helpers.HandleError(err)

	// todo prepend the xml version
	return b
}
