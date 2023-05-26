package services

import (
	"extra_sense/entities"

	"github.com/gorilla/sessions"
)

func CreateUserNumber(numbers_store []*entities.UserNumber) ([]*entities.UserNumber, string) {
	user_number := entities.NewUserNumber(0)
	numbers_store = append(numbers_store, user_number)
	return numbers_store, user_number.ID
}

func FindUserNumber(numbers_store []*entities.UserNumber, userNumberId string) *entities.UserNumber {
	for _, user_number := range numbers_store {
		if user_number.ID == userNumberId {
			return user_number
		}
	}
	return nil
}

func SetUserNumber(session *sessions.Session, userNumberId string, number int) {
	user_numbers := GetStoreOrCreate[entities.UserNumber](session, USER_NUMBERS_STORE_NAME)
	user_number := FindUserNumber(user_numbers, userNumberId)
	user_number.Number = number
	session.Values[USER_NUMBERS_STORE_NAME] = user_numbers
}

func GetUserNumbers(session *sessions.Session) []*entities.UserNumber {
	return GetStoreOrCreate[entities.UserNumber](session, USER_NUMBERS_STORE_NAME)
}
