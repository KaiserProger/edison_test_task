package services

import (
	"errors"
	"extra_sense/dto"
	"extra_sense/entities"
	"extra_sense/globals"
	"math/rand"

	"github.com/gorilla/sessions"
)

func PollExtrasenses(session *sessions.Session, userNumberId string) []dto.GuessTransactionDto {
	transaction_to_return := make([]*entities.GuessTransaction, 0)
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, globals.EXTRASENSES_STORE_NAME)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, globals.TRANSACTIONS_STORE_NAME)
	existing := FindTransactionsByUserNumberId(transactions, userNumberId)
	if len(existing) > 0 {
		dtos := make([]dto.GuessTransactionDto, len(existing))
		for j, transaction := range existing {
			dtos[j] = dto.NewGuessTransactionDto(transaction.GuessedBy, transaction.Number, transaction.UserNumberID)
		}
		return dtos
	}
	for _, extrasense := range extrasenses {
		guess := 10 + (rand.Int() % 90)
		guess_transaction := entities.NewGuessTransaction(extrasense, guess, userNumberId)
		transaction_to_return = append(transaction_to_return, guess_transaction)
	}
	transactions = append(transactions, transaction_to_return...)
	session.Values[globals.TRANSACTIONS_STORE_NAME] = transactions
	dtos := make([]dto.GuessTransactionDto, len(transaction_to_return))
	for j, transaction := range transaction_to_return {
		dtos[j] = dto.NewGuessTransactionDto(transaction.GuessedBy, transaction.Number, transaction.UserNumberID)
	}
	return dtos
}

func CheckExtrasenseAnswers(session *sessions.Session, number int, requestID string) error {
	if number < 10 || number > 99 {
		return errors.New("number doesn't match criteria")
	}
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, globals.EXTRASENSES_STORE_NAME)
	user_numbers := GetStoreOrCreate[entities.UserNumber](session, globals.USER_NUMBERS_STORE_NAME)
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, globals.TRANSACTIONS_STORE_NAME)
	for _, extrasense := range extrasenses {
		transaction := FindTransaction(transactions, extrasense.Name, requestID)
		if transaction == nil {
			continue
		}
		condition := transaction.Number == number
		extrasense.ModifyTrust(condition)
	}
	FindUserNumber(user_numbers, requestID).Number = number
	session.Values[globals.USER_NUMBERS_STORE_NAME] = user_numbers
	session.Values[globals.EXTRASENSES_STORE_NAME] = extrasenses
	return nil
}

func GetExtrasenseTrustLevels(session *sessions.Session) []dto.TrustLevel {
	extrasenses := GetStoreOrCreate[entities.Extrasense](session, globals.EXTRASENSES_STORE_NAME)
	dtos := make([]dto.TrustLevel, len(extrasenses))
	for j, extrasense := range extrasenses {
		dtos[j] = dto.NewTrustLevelDto(extrasense.Name, extrasense.TrustLevel)
	}
	return dtos
}
