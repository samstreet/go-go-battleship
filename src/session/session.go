package session

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	log.Println("session")
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SessionHandler: %v\n", "Session")
}
