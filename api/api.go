package api

import (
	"context"

	account "github.com/NpoolPlatform/message/npool/account/mw/v1"

	"github.com/NpoolPlatform/account-middleware/api/deposit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterManagerServer(server, &Server{})
	deposit.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := account.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
