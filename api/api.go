package api

import (
	"context"

	account1 "github.com/NpoolPlatform/account-middleware/api/account"
	"github.com/NpoolPlatform/account-middleware/api/deposit"
	"github.com/NpoolPlatform/account-middleware/api/goodbenefit"
	"github.com/NpoolPlatform/account-middleware/api/orderbenefit"
	"github.com/NpoolPlatform/account-middleware/api/payment"
	"github.com/NpoolPlatform/account-middleware/api/platform"
	"github.com/NpoolPlatform/account-middleware/api/transfer"
	"github.com/NpoolPlatform/account-middleware/api/user"

	account "github.com/NpoolPlatform/message/npool/account/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterMiddlewareServer(server, &Server{})
	deposit.Register(server)
	account1.Register(server)
	goodbenefit.Register(server)
	payment.Register(server)
	platform.Register(server)
	user.Register(server)
	orderbenefit.Register(server)
	transfer.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := account.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := goodbenefit.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := orderbenefit.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := payment.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := platform.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
