package service

import (
	"context"
	cli "smartcard/internal/app/mongo_client/client"
	"smartcard/internal/app/mongocontrol"
	"smartcard/internal/app/mongocontrol/model"

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

func GetOne(ctx context.Context, objId *primitive.ObjectID) (*model.CardData, error) {
	collection := cli.Mgo.MgoClient.Database("control").Collection("card_data")
	query := bson.M{
		"_id": objId,
	}
	card, err := mongocontrol.GetDataOne(ctx, collection, query, nil)
	if err != nil {
		return nil, err
	}
	return card, nil
}

/* func GetMany(mgo *MgoDriver) {

} */

func AddOne(ctx context.Context, card *model.CardData) (string, error) {
	var result string
	collection := cli.Mgo.MgoClient.Database("control").Collection("card_data")
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
