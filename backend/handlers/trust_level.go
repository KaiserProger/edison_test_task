package handlers

import (
	"encoding/json"
	"extra_sense/fs"
	"extra_sense/globals"
	"extra_sense/services"
	"log"
	"net/http"
)

func GetExtrasenseTrustLevels(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, globals.SESSION_NAME)
	trust_levels := services.GetExtrasenseTrustLevels(session)
	err := json.NewEncoder(writer).Encode(trust_levels)
	if err != nil {
		log.Fatal(err)
	}
}
