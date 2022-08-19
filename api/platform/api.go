package platform

import (
	"github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
	"google.golang.org/grpc"
)

type Server struct {
	platform.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	platform.RegisterMiddlewareServer(server, &Server{})
}
