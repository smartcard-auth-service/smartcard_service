package server

import (
	"net"
	api "smartcard/pkg/grpc/api"
	remote "smartcard/pkg/grpc/grpc_methods"
	log "smartcard/pkg/logging"

	"google.golang.org/grpc"
)

func Run() {
	logger, _ := log.GetLogger()

	s := grpc.NewServer()
	srv := &remote.GRPCServer{}
	api.RegisterScannerSmartCardServer(s, srv)
	logger.Jrn.Println("Register method")

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Jrn.Println(err)
	}

	if err := s.Serve(l); err != nil {
		logger.Jrn.Println(err)
	}
	logger.Jrn.Println("Shutdown")
}

// Я в качестве клиента вызываю методы на получение данных с rust_server
// В качестве сервера регистрирую методы, которые будут отправлять данные на rust_server
