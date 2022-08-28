package service

import (
	"context"
	"smartcard/internal/app/model/mongocontrol"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetOne(ctx context.Context, objId *primitive.ObjectID) (*mongocontrol.CardData, error)
	AddOne(ctx context.Context, card *mongocontrol.CardData) (*primitive.ObjectID, error)
}
