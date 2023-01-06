package carddatadomain

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardDataRepositorier interface {
	GetOne(ctx context.Context, objId primitive.ObjectID) (*CardData, error)
	AddOne(ctx context.Context, card *CardData) (primitive.ObjectID, error)
}

type CardData struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Owner      string             `bson:"owner" json:"owner"`
	TypeOfCard string             `bson:"type_of_card" json:"type_of_card"`
	CVC        string             `bson:"cvc" json:"cvc"`
	Number     string             `bson:"number" json:"number"`
	ExpireDate time.Time          `bson:"date" json:"date"`
}

func (c *CardData) String() string {
	return fmt.Sprintf("Owner: %v\nTypeOfCard: %v\nCVC: %v\nNumber: %v\nExpireDate: %v\n",
		c.Owner, c.TypeOfCard, c.CVC, c.Number, c.ExpireDate)
}
