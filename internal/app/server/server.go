package server

import (
	"context"

	"smartcard/config"
	"smartcard/internal/app/model/client"
	grpcServ "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
)

func Run() {
	var mgoDriver *client.MgoDriver
	ctx := context.Background()
	initServer(ctx, mgoDriver)

	//defer client.Close(mongoConn)
	grpcServ.Run(mgoDriver)

	log.Logger.Jrn.Println("Shutdown server")
}

func initServer(ctx context.Context, mgoDriver *client.MgoDriver) {
	err := log.InitLogger()
	if err != nil {
		Fatal(err)
	}
	config.InitGlobalConfig()
	client.InitMongoConnection(ctx, mgoDriver)
}
