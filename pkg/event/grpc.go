package event

import (
	proto "github.com/a-inacio/hermit-shell/api/github.com/a-inacio/hermit-shell-grpc"
	"github.com/a-inacio/hermit-shell/internal/server"

	"go.uber.org/zap"
)

type Server struct {
	l zap.SugaredLogger
}

func NewServer() *Server {
	instance := Server{
		l: *zap.S(),
	}

	proto.RegisterHermitShellEventServer(server.GetGrprServer(), &instance)

	return &instance
}

func (s *Server) SubscribeEvent(request *proto.SubscribeEventRequest, server proto.HermitShellEvent_SubscribeEventServer) error {
	return nil
}
