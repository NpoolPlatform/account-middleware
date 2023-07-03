package deposit

import (
	"context"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/deposit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	handler, err := deposit1.NewHandler(
		ctx,
		deposit1.WithID(req.ID),
		deposit1.WithAppID(req.AppID),
		deposit1.WithUserID(req.UserID),
		deposit1.WithCoinTypeID(req.CoinTypeID),
		deposit1.WithAccountID(req.AccountID),
		deposit1.WithAddress(req.Address),
		deposit1.WithActive(req.Active),
		deposit1.WithLocked(req.Locked),
		deposit1.WithLockedBy(req.LockedBy),
		deposit1.WithBlocked(req.Blocked),
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
