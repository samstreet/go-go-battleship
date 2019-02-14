package session

import (
	. "../core/helpers"
	. "../core/services"
	. "../session/services"
	. "../session/structs"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionService := NewSessionService()
	userService := NewUserService()

	identifier := r.Header.Get("X-USER-IDENTIFIER")
	user, _ := userService.FindByUUID(identifier)

	createSessionDTO := CreateSessionDTO{Player1: *user}

	session := sessionService.CreateSession(createSessionDTO)
	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.Write(transformDTOToSchema(session, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
}

func JoinSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionService := NewSessionService()
	userService := NewUserService()

	session := sessionService.FindSessionByUUID(vars["session"])

	identifier := r.Header.Get("X-USER-IDENTIFIER")
	player2, _ := userService.FindByUUID(identifier)

	joinSession := JoinSessionDTO{UUID:session.UUID, Player2:player2.GetID()}
	sessionService.JoinSession(joinSession)

	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.WriteHeader(http.StatusAccepted)
}

func ViewSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionService := NewSessionService()
	sessionOut := sessionService.FindSessionByUUID(vars["session"])
	body := transformDTOToSchema(sessionOut, r.Header.Get("Accept"))

	w.Header().Set("Content-Type", r.Header.Get("Accept"))
	w.Write(body)
	w.WriteHeader(http.StatusOK)
}

func transformDTOToSchema(dto interface{}, output string) []byte {
	if output == "application/json" {
		b, err := json.Marshal(dto)
		HandleError(err)
		return b
	}

	b, err := xml.MarshalIndent(dto, "  ", "    ")
	HandleError(err)

	return []byte(xml.Header + string(b))
}
