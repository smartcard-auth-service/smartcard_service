package server

import (
	"context"
	"sync"

	"smartcard/config"
	client "smartcard/internal/app/mongo_client/client"
	grpc_server "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
	"smartcard/pkg/tls/tls_server"
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
	go grpc_server.Run(ctx, servWg)
	log.Logrus.Debug("Running rpc Server")

	servWg.Add(1)
	go tls_server.Run(ctx, servWg)
	log.Logrus.Debug("Running tls Server")

	servWg.Wait()

	log.Logrus.Debug("Shutdown server")
}

func initServer(ctx context.Context) {
	err := log.InitLogger()
	if err != nil {
		log.Logrus.Fatal("Running tls Server")
	}
	config.InitGlobalConfig()
	client.InitMongoConnection(ctx)
}
