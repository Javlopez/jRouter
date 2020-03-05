package main

import (
	"log"
	"net/http"

	"github.com/Javlopez/jrouter"
)

func main() {
	jr := jrouter.New()
	jr.Handle("/some-end-point", HandlerSimpleSample, "GET, POST")
	err := http.ListenAndServe(":8080", jr)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func HandlerSimpleSample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}
