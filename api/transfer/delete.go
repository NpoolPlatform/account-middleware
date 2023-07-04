package transfer

import (
	"context"

	transfer1 "github.com/NpoolPlatform/account-middleware/pkg/mw/transfer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
)

func (s *Server) DeleteTransfer(ctx context.Context, in *npool.DeleteTransferRequest) (*npool.DeleteTransferResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTransfer",
			"In", in,
		)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := transfer1.NewHandler(
		ctx,
		transfer1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTransfer",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTransferResponse{
		Info: info,
	}, nil
}
