package platform

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	accountchecker "github.com/NpoolPlatform/account-manager/api/account"
	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

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
	case accountmgrpb.AccountUsedFor_UserBenefitHot:
	case accountmgrpb.AccountUsedFor_UserBenefitCold:
	case accountmgrpb.AccountUsedFor_PlatformBenefitCold:
	case accountmgrpb.AccountUsedFor_GasProvider:
	case accountmgrpb.AccountUsedFor_PaymentCollector:
	default:
		logger.Sugar().Errorw("validate", "UsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is invalid")
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
