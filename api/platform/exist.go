//nolint:dupl
package platform

import (
	"context"

	platform1 "github.com/NpoolPlatform/account-middleware/pkg/mw/platform"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAccount(ctx context.Context, in *npool.ExistAccountRequest) (*npool.ExistAccountResponse, error) {
	handler, err := platform1.NewHandler(
		ctx,
		platform1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccount",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccount",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistAccountConds(ctx context.Context, in *npool.ExistAccountCondsRequest) (*npool.ExistAccountCondsResponse, error) {
	handler, err := platform1.NewHandler(
		ctx,
		platform1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistAccountConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAccountCondsResponse{
		Info: info,
	}, nil
}
