package goodbenefit

import (
	"context"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/goodbenefit"

	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	"github.com/google/uuid"
)

func (s *Server) UpdateAccount(ctx context.Context, in *npool.UpdateAccountRequest) (*npool.UpdateAccountResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAccount", "ID", in.GetInfo().GetID(), "err", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().AccountID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetAccountID()); err != nil {
			logger.Sugar().Errorw("UpdateAccount", "AccountID", in.GetInfo().GetAccountID(), "err", err)
			return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		exist, err := accountmgrcli.ExistAccount(ctx, in.GetInfo().GetAccountID())
		if err != nil {
			logger.Sugar().Errorw("UpdateAccount", "AccountID", in.GetInfo().GetAccountID(), "err", err)
			return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if !exist {
			logger.Sugar().Errorw("UpdateAccount", "AccountID", in.GetInfo().GetAccountID(), "exist", exist)
			return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, "AccountID is invalid")
		}
	}
	if in.GetInfo().TransactionID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetTransactionID()); err != nil {
			logger.Sugar().Errorw("UpdateAccount", "TransactionID", in.GetInfo().GetTransactionID(), "err", err)
			return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := goodbenefit1.UpdateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAccount", "err", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: info,
	}, nil
}
