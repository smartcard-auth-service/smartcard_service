package mongo_client

import (
	"context"
	"fmt"
	"os"
	"smartcard/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoDriver struct {
	Ctx       context.Context
	MgoClient *mongo.Client
}

var Mgo *MgoDriver

func InitMongoConnection(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.MONGO_URI))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't connect to Mongo: %v", err)
		return err
	}
	Mgo = &MgoDriver{
		Ctx:       ctx,
		MgoClient: client,
	}
	return nil
}
