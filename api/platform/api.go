package platform

import (
	"context"

	"github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	platform.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	platform.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return platform.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
