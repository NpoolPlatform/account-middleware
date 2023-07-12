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
		user1.WithID(req.ID),
		user1.WithAppID(req.AppID),
		user1.WithUserID(req.UserID),
		user1.WithCoinTypeID(req.CoinTypeID),
		user1.WithAccountID(req.AccountID),
		user1.WithAddress(req.Address),
		user1.WithUsedFor(req.UsedFor),
		user1.WithLabels(req.Labels),
		user1.WithActive(req.Active),
		user1.WithLocked(req.Locked),
		user1.WithBlocked(req.Blocked),
		user1.WithMemo(req.Memo),
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
