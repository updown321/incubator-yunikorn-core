package common

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/apache/incubator-yunikorn-core/pkg/log"
	"github.com/apache/incubator-yunikorn-scheduler-interface/lib/go/si"
)

// Defines Non blocking GRPC server interfaces
type NonBlockingGRPCServer interface {
	// Start services at the endpoint
	Start(endpoint string, ss si.SchedulerServer)
	// Waits for the service to stop
	Wait()
	// Stops the service gracefully
	Stop()
	// Stops the service forcefully
	ForceStop()
}

func NewNonBlockingGRPCServer() NonBlockingGRPCServer {
	return &nonBlockingGRPCServer{}
}

// NonBlocking server
type nonBlockingGRPCServer struct {
	wg     sync.WaitGroup
	server *grpc.Server
}

func (s *nonBlockingGRPCServer) Start(endpoint string, ss si.SchedulerServer) {
	s.wg.Add(1)

	go s.serve(endpoint, ss)
}

func (s *nonBlockingGRPCServer) Wait() {
	s.wg.Wait()
}

func (s *nonBlockingGRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *nonBlockingGRPCServer) ForceStop() {
	s.server.Stop()
}

func ParseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("invalid endpoint: %v", ep)
}

// Logging unary interceptor function to log every RPC call
func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Logger().Debug("GPRC call",
		zap.String("method", info.FullMethod))
	log.Logger().Debug("GPRC request",
		zap.String("request", fmt.Sprintf("%+v", req)))
	resp, err := handler(ctx, req)
	if err != nil {
		log.Logger().Debug("GPRC error", zap.Error(err))
	} else {
		log.Logger().Debug("GPRC response",
			zap.String("response", fmt.Sprintf("%+v", resp)))
	}
	return resp, err
}

// Returns unary interceptor that will be used to intercept the execution of a unary RPC on the gRPC server
func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(logGRPC)
}

func (s *nonBlockingGRPCServer) serve(endpoint string, ss si.SchedulerServer) {
	proto, addr, err := ParseEndpoint(endpoint)
	if err != nil {
		log.Logger().Fatal("fatal error", zap.Error(err))
	}

	if proto == "unix" {
		addr = "/" + addr
		if err = os.Remove(addr); err != nil && !os.IsNotExist(err) {
			log.Logger().Fatal("failed to remove unix domain socket",
				zap.String("uds", addr),
				zap.Error(err))
		}
	}

	var listener net.Listener
	listener, err = net.Listen(proto, addr)
	if err != nil {
		log.Logger().Fatal("failed to listen to address",
			zap.Error(err))
	}

	server := grpc.NewServer(withServerUnaryInterceptor())
	s.server = server

	if ss != nil {
		si.RegisterSchedulerServer(server, ss)
	}

	log.Logger().Info("listening for connections",
		zap.String("address", listener.Addr().String()))

	if err = server.Serve(listener); err != nil {
		log.Logger().Fatal("failed to serve", zap.Error(err))
	}
}
