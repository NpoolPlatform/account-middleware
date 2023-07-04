package payment

import (
	"context"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/mw/payment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (s *Server) CreateAccount(ctx context.Context, in *npool.CreateAccountRequest) (*npool.CreateAccountResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAccount",
			"In", in,
		)
		return &npool.CreateAccountResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID),
		payment1.WithCoinTypeID(req.CoinTypeID),
		payment1.WithAccountID(req.AccountID),
		payment1.WithAddress(req.Address),
		payment1.WithActive(req.Active),
		payment1.WithLocked(req.Locked),
		payment1.WithLockedBy(req.LockedBy),
		payment1.WithBlocked(req.Blocked),
		payment1.WithAvailableAt(req.AvailableAt),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAccountResponse{
		Info: info,
	}, nil
}
