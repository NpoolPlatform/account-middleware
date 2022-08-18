package api

import (
	"context"

	"github.com/NpoolPlatform/account-middleware/api/goodbenefit"

	account "github.com/NpoolPlatform/message/npool/account/mw/v1"

	"github.com/NpoolPlatform/account-middleware/api/deposit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterMiddlewareServer(server, &Server{})
	deposit.Register(server)
	goodbenefit.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := account.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
