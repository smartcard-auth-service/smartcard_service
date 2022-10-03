package grpc

import (
	"context"
	"fmt"
	"smartcard/internal/app/mongocontrol/model"
	service "smartcard/internal/app/mongocontrol/service"

	api "smartcard/internal/grpc/api"
	log "smartcard/pkg/logging"
)

func (server *GRPCServer) RegisterCardData(ctx context.Context, req *api.RegistrateRequest) (*api.RegistrateResponse, error) {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()

	var status service.StateOper
	var errText string
	var id string

	regCardData, err := GetDecodingCardData(req.GetRegData())
	if err != nil {
		log.Logrus.Errorf("Error Unmarshal = %v", err)
		errText = fmt.Sprintf("Error Unmarshal card data : %v", err)
		status = service.FAILED
	}

	if !service.IsFailed(status) {
		id, err = service.AddOne(ctx, regCardData)
		if err != nil {
			log.Logrus.Errorf("Error inserting data : %v", err)
			errText = fmt.Sprintf("Error inserting card data : %v", err)
			status = service.FAILED
		}
	}
	response := &api.RegistrateResponse{
		Id:        id,
		Status:    service.ConvertStatus(string(status)),
		ErrorText: errText,
	}
	return response, nil
}

func (server *GRPCServer) GetCardData(ctx context.Context, req *api.GetDataRequest) (*api.GetDataResponse, error) {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()

	var card *model.CardData
	var byteCard []byte
	var status service.StateOper
	var errText string

	idCard, err := GetIdCard(req.GetId())
	if err != nil {
		errText = fmt.Sprintf("Error get ObjectID : %v", err)
		status = service.FAILED
	}
	if !service.IsFailed(status) {
		card, err = service.GetOne(ctx, idCard)
		if err != nil {
			log.Logrus.Errorf("Error get data : %v", err)
			errText = fmt.Sprintf("Error get card data : %v", err)
			status = service.FAILED
		}
	}
	if !service.IsFailed(status) {
		byteCard, err = GetEncodingCardData(card)
		if err != nil {
			log.Logrus.Errorf("Error Marshal = %v", err)
			errText = fmt.Sprintf("Error Marshal bytes : %v", err)
			status = service.FAILED
		}
	}
	response := &api.GetDataResponse{
		Data:      byteCard,
		Status:    service.ConvertStatus(string(status)),
		ErrorText: errText,
	}
	return response, nil
}
