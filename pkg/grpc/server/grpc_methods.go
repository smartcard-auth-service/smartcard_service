package grpc

import (
	"context"
	"fmt"
	"smartcard/internal/app/model/client"

	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"
)

func (s *GRPCServer) RegisterCardData(ctx context.Context, req *api.RegistrateRequest) (*api.RegistrateResponse, error) {
	var status, errText string
	regCardData, err := getDecodingCardData(req)
	if err != nil {
		log.Logger.Jrn.Printf("Error Unmarshal = %v\n", err)
		errText = fmt.Sprintf("Error Unmarshal card data : %v", err)
		status = string(client.FAILED)
	}

	// to do : add implemention method client.AddOne
	err = client.AddOne(regCardData)
	if err != nil {
		log.Logger.Jrn.Printf("Error inserting data : %v", err)
		errText = fmt.Sprintf("Error inserting card data : %v", err)
		status = string(client.FAILED)
	}
	response := &api.RegistrateResponse{
		Status:    status,
		ErrorText: errText,
	}
	return response, nil
}

func (s *GRPCServer) GetCardData(ctx context.Context, req *api.GetDataRequest) (*api.GetDataResponse, error) {
	var errText string
	idCard, err := getIdCard(req)
	card, err := s.mongo.GetOne(idCard)

	byteCard, err := getEncodingCardData(card)
	if err != nil {
		log.Logger.Jrn.Printf("Error Marshal = %v\n", err)
		errText = fmt.Sprintf("Error Marshal bytes : %v", err)
	}
	response := &api.GetDataResponse{
		Data:      byteCard,
		Status:    string(client.SUCCESS),
		ErrorText: errText,
	}
	return response, nil
}
