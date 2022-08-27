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
// to do запихнуть это выше mongocontrol
func GetOne(dbName string, collectionName string, query bson.M, opts *options.FindOptions) {

}

/* func GetMany(mgo *MgoDriver) {

} */

func AddOne(card *CardData) {

}

/* func UpdateOne(mgo *MgoDriver) {

}
*/
