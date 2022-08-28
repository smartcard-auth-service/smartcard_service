package service

import (
	"smartcard/internal/app/model/mongocontrol"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetOne(objId *primitive.ObjectID) (*mongocontrol.CardData, error)
	AddOne(card *mongocontrol.CardData) error
}
