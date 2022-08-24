package mongocontrol

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MgoDriver struct {
	Ctx       context.Context
	MgoClient *mongo.Client
}
