package client

import (
	"context"
	"smartcard/config"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoDriver struct {
	Ctx       context.Context
	MgoClient *mongo.Client
}

func InitMongoConnection(ctx context.Context, mgoDriver *MgoDriver) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MONGO_URI))
	if err != nil {
		log.Logger.Jrn.Fatalf("Connection establishment error %v", err)
	}
	mgoDriver = &MgoDriver{
		Ctx:       ctx,
		MgoClient: client,
	}
}
