package dto

import "extra_sense/entities"

type TransactionHistoryDto struct {
	CorrectNumber     int                   `json:"correct_number"`
	ExtrasenseGuesses []GuessTransactionDto `json:"extrasense_guesses"`
}

func NewTransactionHistoryDto(correctNumber int, transactions []*entities.GuessTransaction) TransactionHistoryDto {
	dtos := make([]GuessTransactionDto, len(transactions))
	for j, entity := range transactions {
		dtos[j] = NewGuessTransactionDto(entity.GuessedBy, entity.Number, entity.UserNumberID)
	}
	return TransactionHistoryDto{
		CorrectNumber:     correctNumber,
		ExtrasenseGuesses: dtos,
	}
}
