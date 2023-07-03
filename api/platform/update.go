package platform

import (
	"context"

	platform1 "github.com/NpoolPlatform/account-middleware/pkg/mw/platform"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

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
	handler, err := platform1.NewHandler(
		ctx,
		platform1.WithID(req.ID),
		platform1.WithBackup(req.Backup),
		platform1.WithActive(req.Active),
		platform1.WithLocked(req.Locked),
		platform1.WithLockedBy(req.LockedBy),
		platform1.WithBlocked(req.Blocked),
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
