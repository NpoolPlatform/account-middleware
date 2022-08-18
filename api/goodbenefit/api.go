package goodbenefit

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
	"google.golang.org/grpc"
)

type Server struct {
	goodbenefit.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodbenefit.RegisterMiddlewareServer(server, &Server{})
}
