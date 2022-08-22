package grpc

import (
	"context"
	api "smartcard/pkg/grpc/api"
)

type GRPCServer struct {
	api.UnimplementedScannerSmartCardServer
}

func (s *GRPCServer) GetSmartCardInfo(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Src: "Response", ErrorText: "ErrorText is nil"}, nil
}
