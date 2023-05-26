package entities

import (
	"encoding/hex"

	"github.com/google/uuid"
)

type UserNumber struct {
	ID     string
	Number int
}

func NewUserNumber(number int) *UserNumber {
	id, _ := uuid.NewRandom()
	return &UserNumber{
		ID:     hex.EncodeToString(id[:]),
		Number: number,
	}
}
