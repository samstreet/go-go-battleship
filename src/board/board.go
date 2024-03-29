package board

import (
	"../board/model"
	"../board/structs"
	"../core/dbal"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func CreateBoardHandler(w http.ResponseWriter, r *http.Request) {
	var DB = dbal.InitialiseConnection()
	DB.Create(&model.BoardModel{})

	var boardModel model.BoardModel
	DB.First(&boardModel, 1)

	session := structs.BoardOutDTO{UUID: uuid.FromStringOrNil(boardModel.ID)}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

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
