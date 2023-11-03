package payment

import (
	"context"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/mw/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (s *Server) LockAccount(ctx context.Context, in *npool.LockAccountRequest) (*npool.LockAccountResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(&in.ID, true),
		payment1.WithLockedBy(&in.LockedBy, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"LockAccount",
			"In", in,
			"Error", err,
		)
		return &npool.LockAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.LockAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"LockAccount",
			"In", in,
			"Error", err,
		)
		return &npool.LockAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockAccountResponse{
		Info: info,
	}, nil
}
