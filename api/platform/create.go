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
		platform1.WithID(req.ID),
		platform1.WithCoinTypeID(req.CoinTypeID),
		platform1.WithUsedFor(req.UsedFor),
		platform1.WithAccountID(req.AccountID),
		platform1.WithAddress(req.Address),
		platform1.WithBackup(req.Backup),
		platform1.WithActive(req.Active),
		platform1.WithLocked(req.Locked),
		platform1.WithLockedBy(req.LockedBy),
		platform1.WithBlocked(req.Blocked),
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
