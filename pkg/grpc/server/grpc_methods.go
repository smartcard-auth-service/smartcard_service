package grpc

import (
	"context"
	"smartcard/internal/app/model/mongocontrol"
	api "smartcard/pkg/grpc/api"
)

func (s *GRPCServer) RegisterCardData(ctx context.Context, req *api.RegistrateRequest) (*api.RegistrateResponse, error) {
	// to do : add implemention method mongocontrol.AddOne
	response := &api.RegistrateResponse{
		Status:    string(mongocontrol.SUCCESS),
		ErrorText: "",
	}
	return response, nil
}

func (s *GRPCServer) GetCardData(ctx context.Context, req *api.GetDataRequest) (*api.GetDataResponse, error) {
	// to do : add implement method mongocontrol.GetOne
	response := &api.GetDataResponse{
		Data:      []byte{}, // dummy
		Status:    string(mongocontrol.SUCCESS),
		ErrorText: "",
	}
	return response, nil
}
