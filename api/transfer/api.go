package transfer

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	transfer.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	transfer.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
