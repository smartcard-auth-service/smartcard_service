package grpc

import (
	"context"
	"net"
	mongo_client "smartcard/internal/app/mongo_client/client"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
	api "smartcard/internal/grpc/api"
	"smartcard/pkg/config"
	log "smartcard/pkg/logging"
	"sync"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	api.UnimplementedScannerSmartCardServer
	cardDataInteractor *carddatacases.CardDataInteractor
}

func Run(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	server := grpc.NewServer()

	grpcServer := &GRPCServer{
		cardDataInteractor: carddatacases.NewCardDataInteractor(mongo_client.Mgo),
	}
	api.RegisterScannerSmartCardServer(server, grpcServer)

	listen, err := net.Listen("tcp", config.Cfg.GRPC_LISTEN_HOST)
	if err != nil {
		log.Logrus.Info("Shutdown application")
		log.Logrus.Fatalf("Error establishing tcp connection on host = %v, error = %v", config.Cfg.GRPC_LISTEN_HOST, err)
	}
	log.Logrus.Info("GRPC Server ready to accept request on host = ", config.Cfg.GRPC_LISTEN_HOST)
	err = server.Serve(listen)
	if err != nil {
		log.Logrus.Info("Shutdown application")
		log.Logrus.Fatal(err)
	}
}
