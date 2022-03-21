package event

import (
	"fmt"
	"sync"
	"time"

	proto "github.com/a-inacio/hermit-shell/api/github.com/a-inacio/hermit-shell-grpc"
	"github.com/a-inacio/hermit-shell/internal/server"

	"go.uber.org/zap"
)

type Server struct {
	l zap.SugaredLogger
}

func RegisterServer() *Server {
	instance := Server{
		l: *zap.S(),
	}

	proto.RegisterHermitShellEventServer(server.GetGrprServer(), &instance)

	return &instance
}

func (s *Server) SubscribeEvent(request *proto.SubscribeEventRequest, server proto.HermitShellEvent_SubscribeEventServer) error {
	log := s.l

	log.Info("SubscribeEvent: ", request.EventName)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			time.Sleep(time.Duration(count) * time.Second)
			resp := proto.SubscribeEventReply{Payload: fmt.Sprintf("Request #%d For Id:%s", count, request.EventName)}
			if err := server.Send(&resp); err != nil {
				log.Errorf("send error %v", err)
			}
			log.Infof("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()

	return nil
}
