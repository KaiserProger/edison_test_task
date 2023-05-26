package handlers

import (
	"extra_sense/fs"
	"extra_sense/services"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CheckNumbers(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, services.SESSION_NAME)
	number, err := strconv.Atoi(request.URL.Query().Get("number"))
	if err != nil {
		log.Print(err)
		fmt.Fprint(writer, http.StatusBadRequest)
		return
	}
	cookies := request.Cookies()
	cookie := &http.Cookie{}
	for _, j := range cookies {
		if j.Name == "X-User-Number-Id" {
			cookie = j
		}
	}
	services.CheckExtrasenseAnswers(session, number, cookie.Value)
	err = session.Save(request, writer)
	if err != nil {
		log.Fatal(err)
	}
}
