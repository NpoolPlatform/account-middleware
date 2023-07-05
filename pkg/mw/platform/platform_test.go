package platform

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"
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

var ret = npool.Account{
	ID:          uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
	AccountID:   uuid.NewString(),
	UsedFor:     basetypes.AccountUsedFor_UserBenefitHot,
	UsedForStr:  basetypes.AccountUsedFor_UserBenefitHot.String(),
	Address:     uuid.NewString(),
	Active:      true,
	LockedByStr: basetypes.AccountLockedBy_DefaultLockedBy.String(),
}

func creatAccount(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCoinTypeID(&ret.CoinTypeID),
		WithAccountID(&ret.AccountID),
		WithAddress(&ret.Address),
		WithUsedFor(&ret.UsedFor),
	)
	assert.Nil(t, err)
	info, err := handler.CreateAccount(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateAccount(t *testing.T) {
	ret.Active = false
	ret.Locked = true
	ret.LockedBy = basetypes.AccountLockedBy_Payment
	ret.LockedByStr = basetypes.AccountLockedBy_Payment.String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithActive(&ret.Active),
		WithLocked(&ret.Locked),
		WithLockedBy(&ret.LockedBy),
		WithBlocked(&ret.Blocked),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	ret.Locked = false
	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithLocked(&ret.Locked),
	)
	assert.Nil(t, err)

	info, err = handler.UpdateAccount(context.Background())
	if assert.Nil(t, err) {
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
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsedFor)},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
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