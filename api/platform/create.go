package platform

import (
	"context"

	platform1 "github.com/NpoolPlatform/account-middleware/pkg/mw/platform"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

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
	handler, err := platform1.NewHandler(
		ctx,
		platform1.WithEntID(req.EntID, false),
		platform1.WithCoinTypeID(req.CoinTypeID, true),
		platform1.WithUsedFor(req.UsedFor, true),
		platform1.WithAccountID(req.AccountID, false),
		platform1.WithAddress(req.Address, true),
		platform1.WithBackup(req.Backup, false),
		platform1.WithActive(req.Active, false),
		platform1.WithLocked(req.Locked, false),
		platform1.WithLockedBy(req.LockedBy, false),
		platform1.WithBlocked(req.Blocked, false),
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
