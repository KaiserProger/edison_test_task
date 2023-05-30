package services

import (
	"extra_sense/dto"
	"extra_sense/entities"
	"extra_sense/globals"

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
	allTransactions := GetStoreOrCreate[entities.GuessTransaction](session, globals.TRANSACTIONS_STORE_NAME)
	transactionMap := make(map[string][]*entities.GuessTransaction, 0)
	for _, j := range allTransactions {
		record, ok := transactionMap[j.UserNumberID]
		if !ok {
			record = make([]*entities.GuessTransaction, 0)
		}
		record = append(record, j)
		transactionMap[j.UserNumberID] = record
	}
	userNumbers := GetUserNumbers(session)
	dtos := make([]dto.TransactionHistoryDto, len(userNumbers))
	for j, userNumber := range userNumbers {
		record, ok := transactionMap[userNumber.ID]
		if !ok {
			continue
		}
		dtos[j] = dto.NewTransactionHistoryDto(userNumber.Number, record)
	}
	return dtos
}
