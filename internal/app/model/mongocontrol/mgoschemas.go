package mongocontrol

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardData struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Owner      string             `bson:"owner" json:"owner"`
	TypeOfCard string             `bson:"type_of_card" json:"type_of_card"`
	CVC        string             `bson:"cvc" json:"cvc"`
	Number     string             `bson:"number" json:"number"`
	ExpireDate time.Time          `bson:"date" json:"date"`
}
