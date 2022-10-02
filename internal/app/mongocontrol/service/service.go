package service

import (
	"context"
	"smartcard/internal/app/mongocontrol/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetOne(ctx context.Context, objId *primitive.ObjectID) (*model.CardData, error)
	AddOne(ctx context.Context, card *model.CardData) (*primitive.ObjectID, error)
}
