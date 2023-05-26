package handlers

import (
	"encoding/json"
	"extra_sense/fs"
	"extra_sense/services"
	"log"
	"net/http"
)

func GetUserNumbers(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, services.SESSION_NAME)
	err := json.NewEncoder(writer).Encode(services.GetUserNumbers(session))
	if err != nil {
		log.Fatal(err)
	}
}
