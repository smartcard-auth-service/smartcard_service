package grpc

import (
	"context"
	"fmt"
	"smartcard/internal/app/model/mongocontrol"
	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"
)

func (s *GRPCServer) RegisterCardData(ctx context.Context, req *api.RegistrateRequest) (*api.RegistrateResponse, error) {
	var errText string
	regCardData, err := getDecodingCardData(ctx, req)
	if err != nil {
		log.Logger.Jrn.Printf("Error Unmarshal = %v\n", err)
		errText = fmt.Sprintf("Error Unmarshal card data : %v", err)
	}

	// to do : add implemention method mongocontrol.AddOne
	mongocontrol.AddOne(regCardData)
	response := &api.RegistrateResponse{
		Status:    string(mongocontrol.SUCCESS),
		ErrorText: errText,
	}
	return response, nil
}

func (s *GRPCServer) GetCardData(ctx context.Context, req *api.GetDataRequest) (*api.GetDataResponse, error) {
	// to do : add implement method mongocontrol.GetOne
	var card *mongocontrol.CardData
	var errText string

	byteCard, err := getEncodingCardData(ctx, card)
	if err != nil {
		log.Logger.Jrn.Printf("Error Marshal = %v\n", err)
		errText = fmt.Sprintf("Error Marshal bytes : %v", err)
	}
	response := &api.GetDataResponse{
		Data:      byteCard,
		Status:    string(mongocontrol.SUCCESS),
		ErrorText: errText,
	}
	return response, nil
}
