package deposit

import (
	"context"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/deposit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SubBalance(ctx context.Context, in *npool.SubBalanceRequest) (*npool.SubBalanceResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"SubBalance",
			"In", in,
		)
		return &npool.SubBalanceResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := deposit1.NewHandler(
		ctx,
		deposit1.WithID(req.ID),
		deposit1.WithIncoming(req.Incoming),
		deposit1.WithOutcoming(req.Outcoming),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SubBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SubBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.SubBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SubBalance",
			"In", in,
			"Error", err,
		)
		return &npool.SubBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SubBalanceResponse{
		Info: info,
	}, nil
}
