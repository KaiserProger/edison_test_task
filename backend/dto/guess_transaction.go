package dto

type GuessTransactionDto struct {
	GuessedBy    string `json:"guessed_by"`
	Number       int    `json:"number"`
	UserNumberID string `json:"user_number_id"`
}

func NewGuessTransactionDto(guessedBy string, number int, userNumberId string) GuessTransactionDto {
	return GuessTransactionDto{
		GuessedBy:    guessedBy,
		Number:       number,
		UserNumberID: userNumberId,
	}
}
