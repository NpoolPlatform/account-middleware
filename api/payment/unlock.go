package payment

import (
	"context"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/mw/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (s *Server) UnlockAccount(ctx context.Context, in *npool.UnlockAccountRequest) (*npool.UnlockAccountResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UnlockAccount",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UnlockAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UnlockAccount",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnlockAccountResponse{
		Info: info,
	}, nil
}
