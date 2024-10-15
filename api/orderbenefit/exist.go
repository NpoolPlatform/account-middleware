//nolint:nolintlint,dupl
package orderbenefit

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	orderbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/orderbenefit"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
)

func (s *Server) ExistAccountConds(ctx context.Context, in *npool.ExistAccountCondsRequest) (*npool.ExistAccountCondsResponse, error) {
	handler, err := orderbenefit1.NewHandler(
		ctx,
		orderbenefit1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	exist, err := handler.ExistAccountConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAccountConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAccountCondsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.ExistAccountCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAccount(ctx context.Context, in *npool.ExistAccountRequest) (*npool.ExistAccountResponse, error) {
	handler, err := orderbenefit1.NewHandler(
		ctx,
		orderbenefit1.WithEntID(&in.EntID, true),
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
