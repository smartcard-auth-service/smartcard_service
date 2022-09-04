package mongo_client

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

var Mgo *MgoDriver

func InitMongoConnection(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MONGO_URI))
	if err != nil {
		log.Logrus.Fatal("Connection establishment error %v", err)
	}
	Mgo = &MgoDriver{
		Ctx:       ctx,
		MgoClient: client,
	}
}
