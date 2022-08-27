package grpc

import (
	"context"
	"encoding/json"
	"smartcard/internal/app/model/mongocontrol"

	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"
)

func getDecodingCardData(ctx context.Context, req *api.RegistrateRequest) (*mongocontrol.CardData, error) {
	var result mongocontrol.CardData
	bytesSlice := req.GetRegData()
	err := json.Unmarshal(bytesSlice, &result)
	if err != nil {
		log.Logger.Jrn.Printf("Error Unmarshal bytes to CardData = %v\n", err)
		return nil, err
	}
	return &result, nil
}

func getEncodingCardData(ctx context.Context, card *mongocontrol.CardData) ([]byte, error) {
	var result []byte
	var err error
	result, err = json.Marshal(card)
	if err != nil {
		log.Logger.Jrn.Printf("Error Marshal CardData to bytes = %v\n", err)
		return nil, err
	}
	return result, nil
}
