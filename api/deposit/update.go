package deposit

import (
	"context"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/deposit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

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
	handler, err := deposit1.NewHandler(
		ctx,
		deposit1.WithID(req.ID),
		deposit1.WithActive(req.Active),
		deposit1.WithLocked(req.Locked),
		deposit1.WithLockedBy(req.LockedBy),
		deposit1.WithBlocked(req.Blocked),
		deposit1.WithCollectingTID(req.CollectingTID),
		deposit1.WithIncoming(req.Incoming),
		deposit1.WithOutcoming(req.Outcoming),
		deposit1.WithScannableAt(req.ScannableAt),
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
