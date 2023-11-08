package goodbenefit

import (
	"context"

	"github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodbenefit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodbenefit.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return goodbenefit.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
