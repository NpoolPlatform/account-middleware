package orderbenefit

import (
	"context"

	orderbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/orderbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
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
	handler, err := orderbenefit1.NewHandler(
		ctx,
		orderbenefit1.WithEntID(req.EntID, false),
		orderbenefit1.WithAppID(req.AppID, true),
		orderbenefit1.WithUserID(req.UserID, true),
		orderbenefit1.WithCoinTypeID(req.CoinTypeID, true),
		orderbenefit1.WithAccountID(req.AccountID, false),
		orderbenefit1.WithOrderID(req.OrderID, true),
		orderbenefit1.WithAddress(req.Address, true),
		orderbenefit1.WithActive(req.Active, false),
		orderbenefit1.WithLocked(req.Locked, false),
		orderbenefit1.WithBlocked(req.Blocked, false),
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
