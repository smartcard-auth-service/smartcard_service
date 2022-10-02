package grpc

import (
	"encoding/json"
	"smartcard/internal/app/mongocontrol/model"

	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDecodingCardData(bytes []byte) (*model.CardData, error) {
	var result model.CardData
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Logrus.Errorf("Error Unmarshal bytes to CardData = %v", err)
		return nil, err
	}
	return &result, nil
}

func getEncodingCardData(card *model.CardData) ([]byte, error) {
	var result []byte
	var err error
	result, err = json.Marshal(card)
	if err != nil {
		log.Logrus.Errorf("Error Marshal CardData to bytes = %v", err)
		return nil, err
	}
	return result, nil
}

func getIdCard(req *api.GetDataRequest) (*primitive.ObjectID, error) {
	var result primitive.ObjectID
	var err error
	log.Logrus.Debugf("reg = %v, type = %T", req.GetId(), req.GetId())
	result, err = primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		log.Logrus.Errorf("Error get ObjectID from request  = %v", err)
		return nil, err
	}
	return &result, nil
}
