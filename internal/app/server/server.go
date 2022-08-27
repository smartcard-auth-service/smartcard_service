package server

import (
	"context"

	"smartcard/internal/app/model/client"
	grpcServ "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
)

func Run() {
	ctx := context.Background()
	err := log.InitLogger()
	if err != nil {
		fatal(err)
	}

	client.InitMongoConnection(ctx)
	//defer client.Close(mongoConn)
	grpcServ.Run()

	log.Logger.Jrn.Println("Shutdown")
}

// Я в качестве клиента вызываю методы на получение данных с rust_server
// В качестве сервера регистрирую методы, которые будут отправлять данные на rust_server
