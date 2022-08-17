package deposit

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	deposit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	deposit.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
