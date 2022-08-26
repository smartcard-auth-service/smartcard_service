package mongocontrol

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetData(mgo *MgoDriver, dbName string, collectionName string, query bson.M, opts *options.FindOptions) {

}

func GetDataMany(mgo *MgoDriver) {

}

func AddData(mgo *MgoDriver) {

}

func UpdateData(mgo *MgoDriver) {

}
