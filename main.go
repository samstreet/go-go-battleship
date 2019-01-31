package main

import (
	"./src/session"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()

	sessionRouter := router.PathPrefix("/session").Subrouter()
	sessionRouter.HandleFunc("", session.CreateSessionHandler).Methods(http.MethodOptions, http.MethodPost)
	sessionRouter.HandleFunc("/join/{session}", session.JoinSessionHandler).Methods(http.MethodOptions, http.MethodPut)
	sessionRouter.HandleFunc("/{session}", session.ViewSessionHandler).Methods(http.MethodOptions, http.MethodGet)
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
