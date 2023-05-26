package services

import (
	"extra_sense/dto"
	"extra_sense/entities"

	"github.com/gorilla/sessions"
)

func FindTransaction(transactions []*entities.GuessTransaction, name string, userNumberId string) *entities.GuessTransaction {
	for _, transaction := range transactions {
		if transaction.GuessedBy == name && transaction.UserNumberID == userNumberId {
			return transaction
		}
	}
	return nil
}

func FindTransactionsByUserNumberId(transactions []*entities.GuessTransaction, userNumberId string) []*entities.GuessTransaction {
	result := make([]*entities.GuessTransaction, 0)
	for _, transaction := range transactions {
		if transaction.UserNumberID == userNumberId {
			result = append(result, transaction)
		}
	}
	return result
}

func GetAllTransactions(session *sessions.Session) []dto.TransactionHistoryDto {
	transactions := GetStoreOrCreate[entities.GuessTransaction](session, TRANSACTIONS_STORE_NAME)
	userNumbers := GetStoreOrCreate[entities.UserNumber](session, USER_NUMBERS_STORE_NAME)
	dtos := make([]dto.TransactionHistoryDto, len(userNumbers))
	for j, userNumber := range userNumbers {
		dtos[j] = dto.NewTransactionHistoryDto(userNumber.Number, FindTransactionsByUserNumberId(transactions, userNumber.ID))
	}
	return dtos
}
