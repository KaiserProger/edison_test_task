package handlers

import (
	"extra_sense/fs"
	"extra_sense/globals"
	"extra_sense/services"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CheckNumbers(writer http.ResponseWriter, request *http.Request) {
	store := fs.NewFSStore()
	session, _ := store.Get(request, globals.SESSION_NAME)
	number, err := strconv.Atoi(request.URL.Query().Get(globals.NUMBER_PARAM_NAME))
	if err != nil {
		log.Print(err)
		fmt.Fprint(writer, http.StatusBadRequest)
		return
	}
	cookies := request.Cookies()
	cookie := &http.Cookie{}
	for _, j := range cookies {
		if j.Name == globals.COOKIE_NUMBER_NAME {
			cookie = j
		}
	}
	err = services.CheckExtrasenseAnswers(session, number, cookie.Value)
	if err != nil {
		fmt.Fprint(writer, http.StatusBadRequest)
	}
	err = session.Save(request, writer)
	if err != nil {
		log.Fatal(err)
	}
}
