package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
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
	ID:            uuid.NewString(),
	GoodID:        uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	AccountID:     uuid.NewString(),
	Address:       uuid.NewString(),
	Backup:        true,
	Active:        true,
	Locked:        false,
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
	Blocked:       false,
	TransactionID: uuid.Nil.String(),
}

var req = npool.AccountReq{
	ID:         &ret.ID,
	GoodID:     &ret.GoodID,
	CoinTypeID: &ret.CoinTypeID,
	AccountID:  &ret.AccountID,
	Address:    &ret.Address,
	Backup:     &ret.Backup,
	Active:     &ret.Active,
	Locked:     &ret.Locked,
	Blocked:    &ret.Blocked,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateAccount(t *testing.T) {
	locked := true
	lockedBy := basetypes.AccountLockedBy_Collecting
	blocked := true
	active := false

	ret.Active = active
	ret.Blocked = blocked
	ret.Locked = locked
	ret.LockedBy = lockedBy
	ret.LockedByStr = lockedBy.String()

	req.Active = &active
	req.Blocked = &blocked
	req.Locked = &locked
	req.LockedBy = &lockedBy

	info, err := UpdateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAccount(t *testing.T) {
	info, err := GetAccount(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAccounts(t *testing.T) {
	infos, total, err := GetAccounts(
		context.Background(),
		&npool.Conds{
			ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
			Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
		},
		0,
		int32(2),
	)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, uint32(1)) {
			assert.Equal(t, infos[0], &ret)
		}
	}
}

func deleteAccount(t *testing.T) {
	info, err := DeleteAccount(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAccount", createAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
