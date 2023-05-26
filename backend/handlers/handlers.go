package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterHandlers(mux *mux.Router) {
	mux.HandleFunc("/", InitGame).Methods("GET")
	mux.HandleFunc("/poll", PollExtrasenses).Methods("GET")
	mux.HandleFunc("/check", CheckNumbers).Methods("GET")
	mux.HandleFunc("/transactions", GetTransactions).Methods("GET")
	mux.HandleFunc("/numbers", GetUserNumbers).Methods("GET")
	mux.HandleFunc("/trust", GetExtrasenseTrustLevels).Methods("GET")
}
