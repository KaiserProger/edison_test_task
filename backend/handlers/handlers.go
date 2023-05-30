package handlers

import (
	"extra_sense/globals"

	"github.com/gorilla/mux"
)

func RegisterHandlers(mux *mux.Router) {
	mux.HandleFunc(globals.INDEX_ROUTE, InitGame).Methods("GET")
	mux.HandleFunc(globals.POLL_EXTRASENSES_ROUTE, PollExtrasenses).Methods("GET")
	mux.HandleFunc(globals.CHECK_EXTRASENSE_ANSWERS_ROUTE, CheckNumbers).Methods("GET")
	mux.HandleFunc(globals.GET_TRANSACTIONS_ROUTE, GetTransactions).Methods("GET")
	mux.HandleFunc(globals.GET_USER_NUMBERS_ROUTE, GetUserNumbers).Methods("GET")
	mux.HandleFunc(globals.GET_EXTRASENSE_TRUST_LEVELS_ROUTE, GetExtrasenseTrustLevels).Methods("GET")
}
