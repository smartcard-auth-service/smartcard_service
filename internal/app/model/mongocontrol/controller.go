package mongocontrol

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StateOper string

var (
	SUCCESS StateOper = "Success"
	FAILED  StateOper = "Failed"
)

func GetOne(mgo *MgoDriver, dbName string, collectionName string, query bson.M, opts *options.FindOptions) {

}

func GetMany(mgo *MgoDriver) {

}

func AddOne(mgo *MgoDriver) {

}

func UpdateOne(mgo *MgoDriver) {

}
