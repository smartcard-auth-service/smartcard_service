package carddatacontroller

import (
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
)

type CardDataController struct {
	cardDataInteractor *carddatacases.CardDataInteractor
}

// Constructor for CardDataController
func NewCardDataController(cardDataInteractor *carddatacases.CardDataInteractor) *CardDataController {
	o := new(CardDataController)
	o.cardDataInteractor = cardDataInteractor
	return o
}
