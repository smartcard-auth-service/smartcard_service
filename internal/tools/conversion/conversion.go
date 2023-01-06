package conversion

import (
	"encoding/json"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDecodingCardData(bytes []byte) (*carddatacases.CardData, error) {
	var result carddatacases.CardData
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Logrus.Errorf("Error Unmarshal bytes to CardData = %v", err)
		return nil, err
	}
	return &result, nil
}

func GetEncodingCardData(card *carddatacases.CardData) ([]byte, error) {
	var result []byte
	var err error
	result, err = json.Marshal(card)
	if err != nil {
		log.Logrus.Errorf("Error Marshal CardData to bytes = %v", err)
		return nil, err
	}
	return result, nil
}

func GetIdCard(id string) (primitive.ObjectID, error) {
	var result primitive.ObjectID
	var err error
	result, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Logrus.Errorf("Error get ObjectID from request  = %v", err)
		return primitive.ObjectID{}, err
	}
	return result, nil
}
