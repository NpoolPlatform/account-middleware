package payment

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
	"google.golang.org/grpc"
)

type Server struct {
	payment.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	payment.RegisterMiddlewareServer(server, &Server{})
}
