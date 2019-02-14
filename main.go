package main

import (
	"./src/board"
	. "./src/board/model"
	. "./src/core/dbal"
	. "./src/core/helpers"
	. "./src/core/model"
	. "./src/core/web/middleware"
	"./src/session"
	. "./src/session/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	HandleError(err)

	db := InitialiseConnection()
	db.DropTableIfExists(&User{}, &BoardModel{}, &SessionModel{}, &BoardPiece{}, &ShipModel{})
	db.AutoMigrate(&User{}, &BoardModel{}, &SessionModel{}, &BoardPiece{}, &ShipModel{})
	db.Model(&User{}).AddForeignKey("session_id", "session(id)", "RESTRICT", "RESTRICT")

	defaultUser := User{ID: "34c77f05-0306-49d4-aa0b-a35fe01a8b18"}
	db.Create(&defaultUser)
	defaultUser2 := User{ID: "56cc1ed2-aeb7-446c-b03a-32385156d54e"}
	db.Create(&defaultUser2)

	carrier := ShipModel{Name: "Carrier", Length: 5}
	battleship := ShipModel{Name: "Battleship", Length: 4}
	cruiser := ShipModel{Name: "Cruiser", Length: 3}
	submarine := ShipModel{Name: "Submarine", Length: 3}
	destroyer := ShipModel{Name: "Destroyer", Length: 2}

	db.Create(&carrier)
	db.Create(&battleship)
	db.Create(&cruiser)
	db.Create(&submarine)
	db.Create(&destroyer)
}

func main() {
	router := mux.NewRouter()
	router.Use(
		RateLimit,
		AcceptableContentTypes,
	)

	sessionRouter := router.PathPrefix("/session").Subrouter()
	sessionRouter.HandleFunc("", session.CreateSessionHandler).Methods(http.MethodOptions, http.MethodPost).Name("battleships.session.create")
	sessionRouter.HandleFunc("/join/{session}", session.JoinSessionHandler).Methods(http.MethodOptions, http.MethodPut).Name("battleships.session.join")
	sessionRouter.HandleFunc("/{session}", session.ViewSessionHandler).Methods(http.MethodOptions, http.MethodGet, http.MethodHead).Name("battleships.session.view")
	sessionRouter.Methods(http.MethodHead)
	sessionRouter.Use(
		HasUserTokenIdentifierHeader,
		HeaderValidUserToken,
	)

	boardRouter := router.PathPrefix("/board").Subrouter()
	boardRouter.HandleFunc("/{uuid}/move", board.UpdateBoardHandler).Methods(http.MethodOptions, http.MethodPut).Name("battleships.board.move")
	boardRouter.Handle("/{uuid}/pieces", SchemaValidator(http.HandlerFunc(board.SetBoardPieces))).Methods(http.MethodOptions, http.MethodPost).Name("battleships.board.pieces")
	boardRouter.Use(
		HasUserTokenIdentifierHeader,
		HeaderValidUserToken,
		BoardExists,
		UserIsAttachedToBoard,
	)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
