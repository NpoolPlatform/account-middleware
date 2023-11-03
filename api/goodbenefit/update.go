package goodbenefit

import (
	"context"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/goodbenefit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID, true),
		goodbenefit1.WithBackup(req.Backup, false),
		goodbenefit1.WithActive(req.Active, false),
		goodbenefit1.WithLocked(req.Locked, false),
		goodbenefit1.WithLockedBy(req.LockedBy, false),
		goodbenefit1.WithBlocked(req.Blocked, false),
		goodbenefit1.WithTransactionID(req.TransactionID, false),
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
