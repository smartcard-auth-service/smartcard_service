package server

import (
	"context"
	"net"

	"smartcard/internal/app/model/mongocontrol/client"
	api "smartcard/pkg/grpc/api"
	remote "smartcard/pkg/grpc/grpc_methods"
	log "smartcard/pkg/logging"

	"google.golang.org/grpc"
)

func Run() {
	ctx := context.Background()
	err := log.InitLogger()
	if err != nil {
		fatal(err)
	}

	mongoConn := client.InitMongoConnection(ctx)
	//defer client.Close(mongoConn)

	server := grpc.NewServer()

	customGrpcServer := &remote.GRPCServer{}
	api.RegisterScannerSmartCardServer(server, customGrpcServer)
	log.Logger.Jrn.Println("Register method")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Logger.Jrn.Println(err)
	}

	if err := server.Serve(listener); err != nil {
		log.Logger.Jrn.Println(err)
	}
	log.Logger.Jrn.Println("Shutdown")
}

// Я в качестве клиента вызываю методы на получение данных с rust_server
// В качестве сервера регистрирую методы, которые будут отправлять данные на rust_server
