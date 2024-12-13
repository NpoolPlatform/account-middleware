package contract

import (
	"context"

	contract1 "github.com/NpoolPlatform/account-middleware/pkg/mw/contract"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"

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

	handler, err := contract1.NewHandler(
		ctx,
		contract1.WithEntID(req.EntID, false),
		contract1.WithGoodID(req.GoodID, true),
		contract1.WithCoinTypeID(req.CoinTypeID, true),
		contract1.WithBackup(req.Backup, false),
		contract1.WithAccountID(req.AccountID, false),
		contract1.WithAddress(req.Address, true),
		contract1.WithActive(req.Active, false),
		contract1.WithLocked(req.Locked, false),
		contract1.WithLockedBy(req.LockedBy, false),
		contract1.WithBlocked(req.Blocked, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	err = handler.CreateAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAccount",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAccountResponse{}, nil
}
