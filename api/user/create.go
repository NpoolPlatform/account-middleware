package user

import (
	"context"

	user1 "github.com/NpoolPlatform/account-middleware/pkg/mw/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
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
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(req.EntID, false),
		user1.WithAppID(req.AppID, true),
		user1.WithUserID(req.UserID, true),
		user1.WithCoinTypeID(req.CoinTypeID, true),
		user1.WithAccountID(req.AccountID, false),
		user1.WithAddress(req.Address, true),
		user1.WithUsedFor(req.UsedFor, true),
		user1.WithLabels(req.Labels, false),
		user1.WithActive(req.Active, false),
		user1.WithLocked(req.Locked, false),
		user1.WithBlocked(req.Blocked, false),
		user1.WithMemo(req.Memo, false),
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
