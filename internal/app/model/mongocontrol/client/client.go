package client

import (
	"context"
	"smartcard/config"
	"smartcard/internal/app/model/mongocontrol"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoConnection(ctx context.Context) *mongocontrol.MgoDriver {
	logger := ctx.Value("logger").(*log.CustomLogger)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		logger.Jrn.Fatalf("Connection establishment error %v", err)
	}
	return &mongocontrol.MgoDriver{
		Ctx:       ctx,
		MgoClient: client,
	}
}
