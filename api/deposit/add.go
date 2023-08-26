package deposit

import (
	"context"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/deposit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddBalance(ctx context.Context, in *npool.AddBalanceRequest) (*npool.AddBalanceResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := deposit1.NewHandler(
		ctx,
		deposit1.WithID(req.ID),
		deposit1.WithIncoming(req.Incoming),
		deposit1.WithOutcoming(req.Outcoming),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
			"Error", err,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.AddBalance(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddBalance",
			"In", in,
			"Error", err,
		)
		return &npool.AddBalanceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddBalanceResponse{
		Info: info,
	}, nil
}
