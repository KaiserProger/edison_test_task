package entities

type GuessTransaction struct {
	GuessedBy    string
	Number       int
	UserNumberID string
}

func NewGuessTransaction(guessedBy *Extrasense, number int, userNumberId string) *GuessTransaction {
	return &GuessTransaction{
		GuessedBy:    guessedBy.Name,
		Number:       number,
		UserNumberID: userNumberId,
	}
}
