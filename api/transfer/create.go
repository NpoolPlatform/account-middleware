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
		transfer1.WithID(req.ID),
		transfer1.WithAppID(req.AppID),
		transfer1.WithUserID(req.UserID),
		transfer1.WithTargetUserID(req.TargetUserID),
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
