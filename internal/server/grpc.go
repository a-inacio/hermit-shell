package server

import (
	"github.com/a-inacio/hermit-shell/internal/logger"
)

func ServeGrpc() {
	log := logger.GetLogger()

	// TODO WIP
	port := 6000
	log.Infow("GRPC server listening", "port", port)
}
