package grpc

import (
	"context"
	"fmt"
	card_data_control "smartcard/internal/app/mongocontrol/card_data_control"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"

	"smartcard/internal/tools/conversion"

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

	var status card_data_control.StateOper
	var errText string
	var id string

	regCardData, err := conversion.GetDecodingCardData(req.GetRegData())
	if err != nil {
		log.Logrus.Errorf("Error Unmarshal = %v", err)
		errText = fmt.Sprintf("Error Unmarshal card data : %v", err)
		status = card_data_control.FAILED
	}

	if !card_data_control.IsFailed(status) {
		rawId, err := server.cardDataInteractor.AddOneCardData(ctx, regCardData)
		if err != nil {
			log.Logrus.Errorf("Error inserting data : %v", err)
			errText = fmt.Sprintf("Error inserting card data : %v", err)
			status = card_data_control.FAILED
		}
		id = rawId.Hex()
	}
	response := &api.RegistrateResponse{
		Id:        id,
		Status:    card_data_control.ConvertStatus(string(status)),
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

	var card *carddatacases.CardData
	var byteCard []byte
	var status card_data_control.StateOper
	var errText string

	idCard, err := conversion.GetIdCard(req.GetId())
	if err != nil {
		errText = fmt.Sprintf("Error get ObjectID : %v", err)
		status = card_data_control.FAILED
	}
	if !card_data_control.IsFailed(status) {
		card, err = server.cardDataInteractor.GetOneCardData(ctx, idCard)
		if err != nil {
			log.Logrus.Errorf("Error get data : %v", err)
			errText = fmt.Sprintf("Error get card data : %v", err)
			status = card_data_control.FAILED
		}
	}
	if !card_data_control.IsFailed(status) {
		byteCard, err = conversion.GetEncodingCardData(card)
		if err != nil {
			log.Logrus.Errorf("Error Marshal = %v", err)
			errText = fmt.Sprintf("Error Marshal bytes : %v", err)
			status = card_data_control.FAILED
		}
	}
	response := &api.GetDataResponse{
		Data:      byteCard,
		Status:    card_data_control.ConvertStatus(string(status)),
		ErrorText: errText,
	}
	return response, nil
}
