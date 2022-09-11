package server

import (
	"context"
	"sync"

	"smartcard/config"
	client "smartcard/internal/app/mongo_client/client"
	grpcServ "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
	tlsSer "smartcard/pkg/tls/tls_server"
)

var servWg sync.WaitGroup

func Run() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()
	ctx := context.Background()
	initServer(ctx)

	servWg.Add(1)
	go grpcServ.Run(ctx, servWg)
	log.Logrus.Debug("Running rpc Server")

	servWg.Add(1)
	go tlsSer.Run(ctx, servWg)
	log.Logrus.Debug("Running tls Server")

	servWg.Wait()

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
