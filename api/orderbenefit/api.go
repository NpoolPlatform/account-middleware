package orderbenefit

import (
	"context"

	"github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	orderbenefit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	orderbenefit.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return orderbenefit.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
