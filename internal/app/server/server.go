package server

import (
	"context"

	"smartcard/config"
	client "smartcard/internal/app/mongo_client"
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
	ctx := context.Background()
	initServer(ctx)

	//defer client.Close(mongoConn)
	grpcServ.Run()

	log.Logrus.Debug("Shutdown server")
}

func initServer(ctx context.Context) {
	err := log.InitLogger()
	if err != nil {
		Fatal(err)
	}
	config.InitGlobalConfig()
	client.InitMongoConnection(ctx)
}
