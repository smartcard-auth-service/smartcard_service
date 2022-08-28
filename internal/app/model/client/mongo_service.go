package client

import (
	"smartcard/internal/app/model/mongocontrol"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StateOper string

var (
	SUCCESS StateOper = "Success"
	FAILED  StateOper = "Failed"
)

func (mgo *MgoDriver) GetOne(objId *primitive.ObjectID) (*mongocontrol.CardData, error) {

}

/* func GetMany(mgo *MgoDriver) {

} */

func AddOne(card *mongocontrol.CardData) error {

}

/* func UpdateOne(mgo *MgoDriver) {

}
*/
