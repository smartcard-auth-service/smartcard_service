package client

import (
	"context"
	"smartcard/config"
	log "smartcard/pkg/logging"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoDriver struct {
	Ctx       context.Context
	MgoClient *mongo.Client
}

var Mgo *MgoDriver
var once sync.Once

func InitMongoConnection(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MONGO_URI))
	if err != nil {
		log.Logger.Jrn.Fatalf("Connection establishment error %v", err)
	}
	once.Do(func() {
		Mgo = &MgoDriver{
			Ctx:       ctx,
			MgoClient: client,
		}
	})
}
