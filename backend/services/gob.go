package services

import (
	"encoding/gob"
	"extra_sense/entities"
)

func RegisterGob() {
	gob.Register([]*entities.Extrasense{})
	gob.Register([]*entities.GuessTransaction{})
	gob.Register([]*entities.UserNumber{})
}
