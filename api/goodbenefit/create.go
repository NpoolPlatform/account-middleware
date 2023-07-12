package goodbenefit

import (
	"context"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/goodbenefit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

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
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID),
		goodbenefit1.WithGoodID(req.GoodID),
		goodbenefit1.WithCoinTypeID(req.CoinTypeID),
		goodbenefit1.WithBackup(req.Backup),
		goodbenefit1.WithAccountID(req.AccountID),
		goodbenefit1.WithAddress(req.Address),
		goodbenefit1.WithActive(req.Active),
		goodbenefit1.WithLocked(req.Locked),
		goodbenefit1.WithLockedBy(req.LockedBy),
		goodbenefit1.WithBlocked(req.Blocked),
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
