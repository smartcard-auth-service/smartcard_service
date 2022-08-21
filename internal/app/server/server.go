package server

import (
	"net"
	remote "smartcard/pkg/grpc"
	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"

	"google.golang.org/grpc"
)

func Run() {
	logger, _ := log.GetLogger()

	s := grpc.NewServer()
	srv := &remote.GRPCServer{}
	api.RegisterScannerSmartCardServer(s, srv)
	logger.Println("Register method")

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Println(err)
	}

	if err := s.Serve(l); err != nil {
		logger.Println(err)
	}
	logger.Println("Shutdown")
}
