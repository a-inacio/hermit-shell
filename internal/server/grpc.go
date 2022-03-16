package server

import (
	"fmt"
	"net"
	"os"

	"github.com/a-inacio/hermit-shell/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var grpcServer *grpc.Server

func GetGrprServer() *grpc.Server {
	return grpcServer
}

func ServeGrpc() {
	log := logger.GetLogger()

	// TODO WIP
	port := "6000"

	log.Infow("GRPC server listening", "port", port)

	grpcServer = grpc.NewServer()

	reflection.Register(grpcServer)

	address := fmt.Sprintf("127.0.0.1:%s", port)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("Unable to create listener", "error", err)

		os.Exit(1)
	}

	log.Info(fmt.Sprintf("Listening on port %s for GRPC requests", address))

	// listen for requests
	_ = grpcServer.Serve(l)
}
