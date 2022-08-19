package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	accountchecker "github.com/NpoolPlatform/account-manager/api/account"
	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(ctx context.Context, info *npool.AccountReq) error {
	usedFor := accountmgrpb.AccountUsedFor_GoodPayment

	if err := accountchecker.Validate(&accountmgrpb.AccountReq{
		CoinTypeID: info.CoinTypeID,
		Address:    info.Address,
		UsedFor:    &usedFor,
	}); err != nil {
		return err
	}

	switch info.GetUsedFor() {
	case accountmgrpb.AccountUsedFor_UserWithdraw:
	case accountmgrpb.AccountUsedFor_UserDeposit:
	case accountmgrpb.AccountUsedFor_UserDirectBenefit:
	default:
		logger.Sugar().Errorw("validate", "UsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	exist, err := accountmgrcli.ExistAccountConds(ctx, &accountmgrpb.Conds{
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GetCoinTypeID(),
		},
		Address: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GetAddress(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "Address", info.GetAddress(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Address is invalid: %v", err))
	}
	if exist {
		logger.Sugar().Errorw("validate", "CoinTypeID", info.GetCoinTypeID(), "Address", info.GetAddress(), "exist", exist)
		return status.Error(codes.AlreadyExists, "Address already exists")
	}

	return nil
}
