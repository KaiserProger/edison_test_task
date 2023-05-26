package handlers

import (
	"encoding/json"
	"extra_sense/fs"
	"extra_sense/services"
	"log"
	"net/http"
)

func PollExtrasenses(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, services.SESSION_NAME)
	cookies := request.Cookies()
	cookie := &http.Cookie{}
	for _, j := range cookies {
		if j.Name == "X-User-Number-Id" {
			cookie = j
		}
	}
	userNumberId := cookie.Value
	transactions := services.PollExtrasenses(session, userNumberId)
	err := session.Save(request, writer)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(writer).Encode(transactions)
	if err != nil {
		log.Fatal(err)
	}
}
