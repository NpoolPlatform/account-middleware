package contract

import (
	"context"

	"github.com/NpoolPlatform/message/npool/account/mw/v1/contract"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	contract.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	contract.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return contract.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
