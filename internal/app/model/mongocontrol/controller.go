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

// to do подумать, как избавиться от зависимости *MgoDriver в этих методах (Вариант с sync.once кажется костыльным)
func GetOne(mgo *MgoDriver, dbName string, collectionName string, query bson.M, opts *options.FindOptions) {

}

func GetMany(mgo *MgoDriver) {

}

func AddOne(mgo *MgoDriver, dbName string, collectionName string, query bson.M, opts *options.FindOptions) {

}

func UpdateOne(mgo *MgoDriver) {

}
