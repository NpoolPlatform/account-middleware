package transfer

import (
	"context"

	transfer1 "github.com/NpoolPlatform/account-middleware/pkg/mw/transfer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
)

func (s *Server) CreateTransfer(ctx context.Context, in *npool.CreateTransferRequest) (*npool.CreateTransferResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTransfer",
			"In", in,
		)
		return &npool.CreateTransferResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithEntID(req.EntID, false),
		transfer1.WithAppID(req.AppID, true),
		transfer1.WithUserID(req.UserID, true),
		transfer1.WithTargetUserID(req.TargetUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTransferResponse{
		Info: info,
	}, nil
}
