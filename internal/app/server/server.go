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
	ctx := getCtx()
	logger := ctx.Value("logger").(*log.CustomLogger)

	mongoConn := client.InitMongoConnection(ctx)
	server := grpc.NewServer()
	customGrpcServer := &remote.GRPCServer{}
	api.RegisterScannerSmartCardServer(server, customGrpcServer)
	logger.Jrn.Println("Register method")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Jrn.Println(err)
	}

	if err := server.Serve(listener); err != nil {
		logger.Jrn.Println(err)
	}
	logger.Jrn.Println("Shutdown")
}

// Я в качестве клиента вызываю методы на получение данных с rust_server
// В качестве сервера регистрирую методы, которые будут отправлять данные на rust_server

func getCtx() context.Context {
	ctx := context.Background()
	logger, _ := log.GetLogger()
	ctx = context.WithValue(ctx, "logger", logger)
	return ctx
}
