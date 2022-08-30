package server

import (
	"context"

	"smartcard/config"
	"smartcard/internal/app/model/client"
	grpcServ "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
)

func Run() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()
	var mgoDriver *client.MgoDriver
	ctx := context.Background()
	initServer(ctx, mgoDriver)

	//defer client.Close(mongoConn)
	grpcServ.Run(mgoDriver)

	log.Logrus.Debug("Shutdown server")
}

func initServer(ctx context.Context, mgoDriver *client.MgoDriver) {
	err := log.InitLogger()
	if err != nil {
		Fatal(err)
	}
	config.InitGlobalConfig()
	client.InitMongoConnection(ctx, mgoDriver)
}
