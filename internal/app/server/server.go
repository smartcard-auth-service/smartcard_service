package server

import (
	"context"
	"fmt"
	"os"
	"sync"

	client "smartcard/internal/app/mongo_client/client"
	grpc_server "smartcard/internal/grpc/server"
	"smartcard/internal/tls/tls_server"
	"smartcard/pkg/config"
	log "smartcard/pkg/logging"
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
	err := initServer(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing server: %v", err)
		os.Exit(1)
	}

	servWg.Add(1)
	go grpc_server.Run(ctx, &servWg)
	log.Logrus.Trace("Running rpc Server")

	servWg.Add(1)
	go tls_server.Run(ctx, &servWg)
	log.Logrus.Trace("Running tls Server")

	servWg.Wait()

	log.Logrus.Trace("Shutdown server")
}

func initServer(ctx context.Context) error {
	config.InitGlobalConfig()
	err := log.InitLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error init logger: %v", err)
		return err
	}
	err = client.InitMongoConnection(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error init mongo connection: %v", err)
		return err
	}
	return nil
}
