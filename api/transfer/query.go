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

func (s *Server) GetTransfer(ctx context.Context, in *npool.GetTransferRequest) (*npool.GetTransferResponse, error) {
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.GetTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.GetTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTransferResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTransfers(ctx context.Context, in *npool.GetTransfersRequest) (*npool.GetTransfersResponse, error) {
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithConds(in.GetConds()),
		transfer1.WithOffset(in.GetOffset()),
		transfer1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTransfers",
			"In", in,
			"Error", err,
		)
		return &npool.GetTransfersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTransfers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTransfers",
			"In", in,
			"Error", err,
		)
		return &npool.GetTransfersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTransfersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
