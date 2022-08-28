package client

import (
	"context"
	"smartcard/internal/app/model/mongocontrol"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StateOper string

var (
	SUCCESS StateOper = "Success"
	FAILED  StateOper = "Failed"
)

func ConvertStatus(status string) string {
	if status == "" {
		return string(SUCCESS)
	}
	return string(FAILED)
}

func (mgo *MgoDriver) GetOne(ctx context.Context, objId *primitive.ObjectID) (*mongocontrol.CardData, error) {
	collection := mgo.MgoClient.Database("control").Collection("card_data")
	query := bson.M{
		"$exists": true,
		"_id":     objId,
	}
	card, err := mongocontrol.GetDataOne(ctx, collection, query, nil)
	if err != nil {
		return nil, err
	}
	return card, nil
}

/* func GetMany(mgo *MgoDriver) {

} */

func (mgo *MgoDriver) AddOne(ctx context.Context, card *mongocontrol.CardData) (string, error) {
	var result string
	collection := mgo.MgoClient.Database("control").Collection("card_data")
	document := bson.D{
		{Key: "type_of_card", Value: card.TypeOfCard},
		{Key: "owner", Value: card.Owner},
		{Key: "cvc", Value: card.CVC},
		{Key: "number", Value: card.Number},
		{Key: "date", Value: card.ExpireDate},
	}
	resultInsert, err := mongocontrol.InsertOne(ctx, collection, document, nil)
	if err != nil {
		return "", err
	}
	result = resultInsert.InsertedID.(primitive.ObjectID).Hex()
	return result, nil
}

/* func UpdateOne(mgo *MgoDriver) {

}
*/
