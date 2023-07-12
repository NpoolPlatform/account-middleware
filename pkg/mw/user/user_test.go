package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Account{
		ID:         uuid.NewString(),
		AppID:      uuid.NewString(),
		UserID:     uuid.NewString(),
		CoinTypeID: uuid.NewString(),
		AccountID:  uuid.NewString(),
		Address:    uuid.NewString(),
		Active:     true,
		Blocked:    false,
		UsedFor:    basetypes.AccountUsedFor_UserWithdraw,
		UsedForStr: basetypes.AccountUsedFor_UserWithdraw.String(),
		Memo:       uuid.NewString(),
		Labels:     []string{uuid.NewString(), uuid.NewString()},
	}
	locked = false
	retReq = npool.AccountReq{
		ID:         &ret.ID,
		AppID:      &ret.AppID,
		UserID:     &ret.UserID,
		CoinTypeID: &ret.CoinTypeID,
		AccountID:  &ret.AccountID,
		Address:    &ret.Address,
		UsedFor:    &ret.UsedFor,
		Labels:     ret.Labels,
		Active:     &ret.Active,
		Blocked:    &ret.Blocked,
		Locked:     &locked,
		Memo:       &ret.Memo,
	}
)

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(retReq.ID),
		WithAppID(retReq.AppID),
		WithUserID(retReq.UserID),
		WithCoinTypeID(retReq.CoinTypeID),
		WithAccountID(retReq.AccountID),
		WithAddress(retReq.Address),
		WithUsedFor(retReq.UsedFor),
		WithLabels(retReq.Labels),
		WithMemo(retReq.Memo),
	)
	assert.Nil(t, err)
	info, err := handler.CreateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.LabelsStr = info.LabelsStr
		assert.Equal(t, info, &ret)
	}
}

func updateAccount(t *testing.T) {
	ret.Active = false
	ret.Labels = []string{uuid.NewString(), uuid.NewString()}
	ret.Blocked = true
	ret.Memo = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(retReq.ID),
		WithActive(&ret.Active),
		WithBlocked(&ret.Blocked),
		WithLabels(ret.Labels),
		WithMemo(&ret.Memo),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.LabelsStr = info.LabelsStr
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)
	info, err := handler.GetAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAccounts(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
		}),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	infos, _, err := handler.GetAccounts(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)
	info, err := handler.DeleteAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAccount(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createAccount", creatAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
