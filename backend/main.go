package main

import (
	"extra_sense/handlers"
	"extra_sense/middlewares"
	"extra_sense/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handler := mux.NewRouter()
	handlers.RegisterHandlers(handler)
	services.RegisterGob()
	server := http.Server{
		Addr:    ":8000",
		Handler: middlewares.Cors(handler),
	}
	log.Fatal(server.ListenAndServe())
}
