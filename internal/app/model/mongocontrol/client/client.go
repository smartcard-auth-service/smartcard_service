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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MONGO_URI))
	if err != nil {
		log.Logger.Jrn.Fatalf("Connection establishment error %v", err)
	}
	return &mongocontrol.MgoDriver{
		Ctx:       ctx,
		MgoClient: client,
	}
}
