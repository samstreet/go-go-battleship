package board

import (
	. "../board/structs"
	. "../core/helpers"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func UpdateBoardHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	session := BoardOutDTO{UUID: id}

	w.Write(TransformDTOToSchema(session, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
}

func SetBoardPieces(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	session := BoardOutDTO{UUID: id}
	w.Write(TransformDTOToSchema(session, r.Header.Get("Accept")))
	w.WriteHeader(http.StatusOK)
}
