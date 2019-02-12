package main

import (
	"./src/board"
	BoardModels "./src/board/model"
	"./src/core/dbal"
	"./src/core/helpers"
	CoreModels "./src/core/model"
	"./src/core/web/middleware"
	"./src/session"
	SessionModels "./src/session/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	helpers.HandleError(err)

	db := dbal.InitialiseConnection()
	db.DropTableIfExists(&CoreModels.User{}, &BoardModels.BoardModel{}, &SessionModels.SessionModel{}, &BoardModels.BoardPiece{}, &BoardModels.ShipModel{})
	db.AutoMigrate(&CoreModels.User{}, &BoardModels.BoardModel{}, &SessionModels.SessionModel{}, &BoardModels.BoardPiece{}, &BoardModels.ShipModel{})
	db.Model(&CoreModels.User{}).AddForeignKey("session_id", "session(id)", "RESTRICT", "RESTRICT")

	defaultUser := CoreModels.User{ID: "34c77f05-0306-49d4-aa0b-a35fe01a8b18"}
	db.Create(&defaultUser)
	defaultUser2 := CoreModels.User{ID: "56cc1ed2-aeb7-446c-b03a-32385156d54e"}
	db.Create(&defaultUser2)

	carrier := BoardModels.ShipModel{Name: "Carrier", Length: 5}
	battleship := BoardModels.ShipModel{Name: "Battleship", Length: 4}
	cruiser := BoardModels.ShipModel{Name: "Cruiser", Length: 3}
	submarine := BoardModels.ShipModel{Name: "Submarine", Length: 3}
	destroyer := BoardModels.ShipModel{Name: "Destroyer", Length: 2}

	db.Create(&carrier)
	db.Create(&battleship)
	db.Create(&cruiser)
	db.Create(&submarine)
	db.Create(&destroyer)
}

func main() {
	router := mux.NewRouter()
	router.Use(
		middleware.RateLimit,
		middleware.AcceptableContentTypes,
	)

	sessionRouter := router.PathPrefix("/session").Subrouter()
	sessionRouter.HandleFunc("", session.CreateSessionHandler).Methods(http.MethodOptions, http.MethodPost)
	sessionRouter.HandleFunc("/join/{session}", session.JoinSessionHandler).Methods(http.MethodOptions, http.MethodPut)
	sessionRouter.HandleFunc("/{session}", session.ViewSessionHandler).Methods(http.MethodOptions, http.MethodGet, http.MethodHead)
	sessionRouter.Methods(http.MethodHead)
	sessionRouter.Use(
		middleware.HasUserTokenIdentifierHeader,
		middleware.HeaderValidUserToken,
	)

	boardRouter := router.PathPrefix("/board").Subrouter()
	boardRouter.HandleFunc("/{uuid}/move", board.UpdateBoardHandler).Methods(http.MethodOptions, http.MethodPut)
	boardRouter.HandleFunc("/{uuid}/pieces", board.SetBoardPieces).Methods(http.MethodOptions, http.MethodPost)
	boardRouter.Use(
		middleware.HasUserTokenIdentifierHeader,
		middleware.HeaderValidUserToken,
		middleware.BoardExists,
		middleware.UserIsAttachedToBoard,
	)

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
