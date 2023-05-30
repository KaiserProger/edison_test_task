package services

import (
	"encoding/hex"
	"extra_sense/entities"
	"extra_sense/globals"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func Populate(session *sessions.Session) string {
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, globals.EXTRASENSES_STORE_NAME)
	// Cleanup user numbers
	user_numbers := GetUserNumbers(session)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, globals.TRANSACTIONS_STORE_NAME)
	user_numbers, userNumberId := CreateUserNumber(user_numbers)
	session.Values[globals.USER_NUMBERS_STORE_NAME] = user_numbers
	session.Values[globals.TRANSACTIONS_STORE_NAME] = transactions
	if len(extrasenses) > 0 {
		return userNumberId
	}
	for j := 0; j < globals.EXTRASENSES_PER_SESSION; j++ {
		id, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}
		extrasenses = append(extrasenses, entities.NewExtrasense(hex.EncodeToString(id[:])))
	}
	session.Values[globals.EXTRASENSES_STORE_NAME] = extrasenses
	return userNumberId
}
