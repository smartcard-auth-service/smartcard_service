package mongocontrol

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardData struct {
	ID         primitive.ObjectID
	Owner      string
	TypeOfCard string
	CSV        string
	Number     string
	ExpireDate time.Time
}
