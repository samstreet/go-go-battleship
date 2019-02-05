package session

import (
	"../core/helpers"
	"../session/services"
	SessionDTO "../session/structs"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionService := services.NewSessionService()
	session := sessionService.CreateSession()
	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.Write(transformDTOToSchema(session, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
}

func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	// todo - build out the logic for joining sessions
	w.WriteHeader(http.StatusAccepted)
}

func ViewSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionService := services.NewSessionService()
	sessionOut := sessionService.FindSessionByUUID(vars["session"])
	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.Write(transformDTOToSchema(sessionOut, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
}

func transformDTOToSchema(dto SessionDTO.SessionOutDTO, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(dto)
		helpers.HandleError(err)
		return b
	}

	b, err := xml.MarshalIndent(dto, "  ", "    ")
	helpers.HandleError(err)

	return []byte(xml.Header + string(b))
}
