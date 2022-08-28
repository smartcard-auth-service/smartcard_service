package grpc

import (
	"encoding/json"
	"smartcard/internal/app/model/mongocontrol"

	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getDecodingCardData(req *api.RegistrateRequest) (*mongocontrol.CardData, error) {
	var result mongocontrol.CardData
	bytesSlice := req.GetRegData()
	err := json.Unmarshal(bytesSlice, &result)
	if err != nil {
		log.Logger.Jrn.Printf("Error Unmarshal bytes to CardData = %v\n", err)
		return nil, err
	}
	return &result, nil
}

func getEncodingCardData(card *mongocontrol.CardData) ([]byte, error) {
	var result []byte
	var err error
	result, err = json.Marshal(card)
	if err != nil {
		log.Logger.Jrn.Printf("Error Marshal CardData to bytes = %v\n", err)
		return nil, err
	}
	return result, nil
}

func getIdCard(req *api.GetDataRequest) (*primitive.ObjectID, error) {
	var result *primitive.ObjectID
	var err error
	bytesSlice := req.GetId()
	err = json.Unmarshal(bytesSlice, &result)
	if err != nil {
		log.Logger.Jrn.Printf("Error Unmarshal bytes to ObjectID = %v\n", err)
		return nil, err
	}
	return result, nil
}
