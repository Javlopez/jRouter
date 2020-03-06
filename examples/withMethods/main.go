package main

import (
	"log"
	"net/http"

	"github.com/Javlopez/jrouter"
)

func main() {
	jr := jrouter.New()
	jr.Get("/users", HandlerSample)
	jr.Post("/users/", HandlerSample)
	jr.Put("/users/", HandlerSample)
	jr.Patch("/users/", HandlerSample)
	err := http.ListenAndServe(":8080", jr)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func HandlerSample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.Method))
}
