package main

import (
	"./src/session"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(os.Getenv("API_VERSION")+"/battleships/game", session.CreateSessionHandler)
	http.Handle("/", r)
}
