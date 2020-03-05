package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Javlopez/jrouter"
)

func main() {
	jr := jrouter.New()
	jr.Handle("/users/{id}", HandlerSimpleSample, "GET")
	err := http.ListenAndServe(":8080", jr)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func HandlerSimpleSample(w http.ResponseWriter, r *http.Request) {
	userid := jrouter.Read(r, "id")
	message := fmt.Sprintf("Your id is %s", userid)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
