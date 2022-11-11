package account

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/account"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
