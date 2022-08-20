package user

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/user"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterMiddlewareServer(server, &Server{})
}
