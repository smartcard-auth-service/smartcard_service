package mongocontrol

import (
	"context"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDataOne(ctx context.Context, collection *mongo.Collection, query bson.M, opts *options.FindOneOptions) (*CardData, error) {
	var result CardData
	err := collection.FindOne(ctx, query, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Logrus.Error("No objects in db which match current query = %v", query)
			return nil, err
		}
		log.Logrus.Error("Error getting object = %v", err)
	}
	return &result, err
}

func GetDataMany(ctx context.Context, collection *mongo.Collection, query bson.M, opts *options.FindOptions) ([]*CardData, error) {
	var result []*CardData
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		log.Logrus.Errorf("Error getting object array = %v", err)
		return nil, err
	}
	for cursor.Next(ctx) {
		var singleResult *CardData
		if err = cursor.Decode(&singleResult); err != nil {
			log.Logrus.Errorf("Error getting cursor object = %v", err)
			return nil, err
		}
		result = append(result, singleResult)
	}
	return result, nil
}

func InsertOne(ctx context.Context, collection *mongo.Collection, document bson.D, opts *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, document, opts)
	if err != nil {
		log.Logrus.Errorf("Error InsertOne = %v", err)
		return nil, err
	}
	return result, nil
}

/* func UpdateOne() {

}

func UpdateMany() {

}


func InsertMany() {

}
*/
