//nolint:nolintlint,dupl
package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	user1 "github.com/NpoolPlatform/account-middleware/pkg/mw/user"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
)

func (s *Server) ExistAccountConds(ctx context.Context, in *npool.ExistAccountCondsRequest) (*npool.ExistAccountCondsResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistAccountConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAccountCondsResponse{
		Info: exist,
	}, nil
}
