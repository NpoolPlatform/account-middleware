package payment

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	accountchecker "github.com/NpoolPlatform/account-manager/api/account"
	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
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

	if info.CollectingTID == nil {
		logger.Sugar().Errorw("validate", "CollectingTID", info.CollectingTID)
		return status.Error(codes.InvalidArgument, "CollectingTID is empty")
	}

	if _, err := uuid.Parse(info.GetCollectingTID()); err != nil {
		logger.Sugar().Errorw("validate", "CollectingTID", info.GetCollectingTID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CollectingTID is invalid: %v", err))
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
