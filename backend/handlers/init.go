package handlers

import (
	"extra_sense/fs"
	"extra_sense/globals"
	"extra_sense/services"
	"log"
	"net/http"
)

func InitGame(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, globals.SESSION_NAME)
	userNumberId := services.Populate(session)
	err := session.Save(request, writer)
	if err != nil {
		log.Fatal(err)
	}
	http.SetCookie(writer, &http.Cookie{
		Name:  globals.COOKIE_NUMBER_NAME,
		Value: userNumberId,
	})
}
