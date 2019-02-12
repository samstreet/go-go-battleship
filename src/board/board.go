package board

import (
	"../board/structs"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func UpdateBoardHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	session := structs.BoardOutDTO{UUID: id}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func SetBoardPieces(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	session := structs.BoardOutDTO{UUID: id}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
