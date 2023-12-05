package payment

import (
	"context"

	"github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	payment.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	payment.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return payment.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
