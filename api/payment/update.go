package payment

import (
	"context"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/mw/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (s *Server) UpdateAccount(ctx context.Context, in *npool.UpdateAccountRequest) (*npool.UpdateAccountResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAccount",
			"In", in,
		)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID, true),
		payment1.WithAccountID(req.AccountID, false),
		payment1.WithActive(req.Active, false),
		payment1.WithLocked(req.Locked, false),
		payment1.WithLockedBy(req.LockedBy, false),
		payment1.WithBlocked(req.Blocked, false),
		payment1.WithCollectingTID(req.CollectingTID, false),
		payment1.WithAvailableAt(req.AvailableAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: info,
	}, nil
}
