package deposit

import (
	"context"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"github.com/google/uuid"
)

func (s *Server) UpdateAccount(
	ctx context.Context, in *npool.UpdateAccountRequest,
) (
	*npool.UpdateAccountResponse, error,
) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAccount", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().CollectingTID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetCollectingTID()); err != nil {
			logger.Sugar().Errorw("UpdateAccount", "CollectingTID", in.GetInfo().GetCollectingTID(), "error", err)
			return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := deposit1.UpdateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAccount", "err", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: info,
	}, nil
}
