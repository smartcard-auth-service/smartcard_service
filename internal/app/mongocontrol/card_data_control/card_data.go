package carddatacontrol

import (
	"context"
	cli "smartcard/internal/app/mongo_client/client"
	carddatadomain "smartcard/internal/clearcore/domain/card_data_domain"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardDataRepository struct {
	mgoDriver *cli.MgoDriver
}

func NewCardDataRepository(mgoDriver *cli.MgoDriver) *CardDataRepository {
	return &CardDataRepository{
		mgoDriver: mgoDriver,
	}
}

type StateOper string

var (
	SUCCESS StateOper = "Success"
	FAILED  StateOper = "Failed"
)

func IsFailed(status StateOper) bool {
	return status == FAILED
}

func ConvertStatus(status string) string {
	if status == "" {
		return string(SUCCESS)
	}
	return string(FAILED)
}

func (rep *CardDataRepository) getCollection() *mongo.Collection {
	return rep.mgoDriver.MgoClient.Database("control").Collection("card_data")
}

func (rep *CardDataRepository) GetOne(ctx context.Context, objId primitive.ObjectID) (*carddatadomain.CardData, error) {
	collection := rep.getCollection()
	query := bson.M{
		"_id": objId,
	}
	var result carddatadomain.CardData
	err := collection.FindOne(ctx, query, nil).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Logrus.Error("No objects in db which match current query = %v", query)
			return nil, err
		}
		log.Logrus.Error("Error getting object = %v", err)
	}
	return &result, nil
}

/* func GetMany(mgo *MgoDriver) {

} */

func (rep *CardDataRepository) AddOne(ctx context.Context, card *carddatadomain.CardData) (primitive.ObjectID, error) {
	collection := rep.getCollection()
	resultInsert, err := collection.InsertOne(ctx, card, nil)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	//result = resultInsert.InsertedID.(primitive.ObjectID).Hex()
	return resultInsert.InsertedID.(primitive.ObjectID), nil
}

/* func UpdateOne(mgo *MgoDriver) {

}
*/
