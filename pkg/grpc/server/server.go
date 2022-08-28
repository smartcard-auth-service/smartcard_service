package grpc

import (
	"net"
	"smartcard/config"
	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"

	"smartcard/internal/app/model/client"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	mongo *client.MgoDriver
	api.UnimplementedScannerSmartCardServer
}

func Run(mgo *client.MgoDriver) {
	server := grpc.NewServer()
	api.RegisterScannerSmartCardServer(server, &GRPCServer{mongo: mgo})

	listen, err := net.Listen("tcp", config.Cfg.GRPC_LISTEN_PORT)
	if err != nil {
		log.Logger.Jrn.Fatalf("Error establishing tcp connection on port = %v, error = %v", config.Cfg.GRPC_LISTEN_PORT, err)
	}
	log.Logger.Jrn.Println("GRPC Server ready to accept request")
	server.Serve(listen)
}
