package services

import (
	"encoding/hex"
	"extra_sense/dto"
	"extra_sense/entities"
	"log"
	"math/rand"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

func Populate(session *sessions.Session) string {
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, EXTRASENSES_STORE_NAME)
	user_numbers := GetStoreOrCreate[entities.UserNumber](session, USER_NUMBERS_STORE_NAME)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, TRANSACTIONS_STORE_NAME)
	user_numbers, userNumberId := CreateUserNumber(user_numbers)
	session.Values[USER_NUMBERS_STORE_NAME] = user_numbers
	session.Values[TRANSACTIONS_STORE_NAME] = transactions
	if len(extrasenses) > 0 {
		return userNumberId
	}
	for j := 0; j < EXTRASENSES_PER_SESSION; j++ {
		id, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}
		extrasenses = append(extrasenses, entities.NewExtrasense(hex.EncodeToString(id[:])))
	}
	session.Values[EXTRASENSES_STORE_NAME] = extrasenses
	return userNumberId
}

func PollExtrasenses(session *sessions.Session, userNumberId string) []dto.GuessTransactionDto {
	transaction_to_return := make([]*entities.GuessTransaction, 0)
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, EXTRASENSES_STORE_NAME)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, TRANSACTIONS_STORE_NAME)
	existing := FindTransactionsByUserNumberId(transactions, userNumberId)
	if len(existing) > 0 {
		dtos := make([]dto.GuessTransactionDto, len(existing))
		for j, transaction := range existing {
			dtos[j] = dto.NewGuessTransactionDto(transaction.GuessedBy, transaction.Number, transaction.UserNumberID)
		}
		return dtos
	}
	for _, extrasense := range extrasenses {
		guess := rand.Int() % 100
		guess_transaction := entities.NewGuessTransaction(extrasense, guess, userNumberId)
		transaction_to_return = append(transaction_to_return, guess_transaction)
	}
	transactions = append(transactions, transaction_to_return...)
	session.Values[TRANSACTIONS_STORE_NAME] = transactions
	dtos := make([]dto.GuessTransactionDto, len(transaction_to_return))
	for j, transaction := range transaction_to_return {
		dtos[j] = dto.NewGuessTransactionDto(transaction.GuessedBy, transaction.Number, transaction.UserNumberID)
	}
	return dtos
}

func CheckExtrasenseAnswers(session *sessions.Session, number int, requestID string) {
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, EXTRASENSES_STORE_NAME)
	user_numbers := GetStoreOrCreate[entities.UserNumber](session, USER_NUMBERS_STORE_NAME)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, TRANSACTIONS_STORE_NAME)
	for _, extrasense := range extrasenses {
		transaction := FindTransaction(transactions, extrasense.Name, requestID)
		if transaction == nil {
			log.Print("No transaction found for extrasense " + extrasense.Name)
			continue
		}
		condition := transaction.Number == number
		extrasense.ModifyTrust(condition)
	}
	FindUserNumber(user_numbers, requestID).Number = number
	session.Values[USER_NUMBERS_STORE_NAME] = user_numbers
	session.Values[EXTRASENSES_STORE_NAME] = extrasenses
}

func GetExtrasenseTrustLevels(session *sessions.Session) []dto.TrustLevel {
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, EXTRASENSES_STORE_NAME)
	dtos := make([]dto.TrustLevel, len(extrasenses))
	for j, extrasense := range extrasenses {
		dtos[j] = dto.NewTrustLevelDto(extrasense.Name, extrasense.TrustLevel)
	}
	return dtos
}
