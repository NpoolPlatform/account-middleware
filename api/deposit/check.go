package deposit

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	accountchecker "github.com/NpoolPlatform/account-manager/api/account"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func validate(info *npool.AccountReq) error {
	usedFor := accountmgrpb.AccountUsedFor_UserDeposit
	platformHoldPrivateKey := true

	if err := accountchecker.Validate(&accountmgrpb.AccountReq{
		CoinTypeID:             info.CoinTypeID,
		Address:                info.Address,
		UsedFor:                &usedFor,
		PlatformHoldPrivateKey: &platformHoldPrivateKey,
	}); err != nil {
		return err
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

	return nil
}
