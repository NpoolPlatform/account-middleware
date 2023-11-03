//nolint:nolintlint,dupl
package transfer

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	transfer1 "github.com/NpoolPlatform/account-middleware/pkg/mw/transfer"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
)

func (s *Server) ExistTransfer(ctx context.Context, in *npool.ExistTransferRequest) (*npool.ExistTransferResponse, error) {
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTransferResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTransferConds(ctx context.Context, in *npool.ExistTransferCondsRequest) (*npool.ExistTransferCondsResponse, error) {
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTransferConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTransferConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTransferConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTransferCondsResponse{
		Info: exist,
	}, nil
}
