package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
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

var ret = &npool.Account{
	ID:            uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	Address:       uuid.NewString(),
	Active:        true,
	Locked:        false,
	LockedBy:      basetypes.AccountLockedBy_DefaultLockedBy,
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
	Blocked:       false,
	CollectingTID: uuid.UUID{}.String(),
}

var retReq = &npool.AccountReq{
	ID:            &ret.ID,
	CoinTypeID:    &ret.CoinTypeID,
	Address:       &ret.Address,
	Active:        &ret.Active,
	Locked:        &ret.Locked,
	LockedBy:      &ret.LockedBy,
	Blocked:       &ret.Blocked,
	CollectingTID: &ret.CollectingTID,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.AccountID = info.AccountID
		ret.AvailableAt = info.AvailableAt
		assert.Equal(t, ret, info)
	}
}

func updateAccount(t *testing.T) {
	active := false
	locked := true
	blocked := true
	collectingTID := uuid.NewString()

	ret.Active = active
	ret.Locked = locked
	ret.Blocked = blocked
	ret.CollectingTID = collectingTID

	retReq.Active = &active
	retReq.Locked = &locked
	retReq.Blocked = &blocked
	retReq.CollectingTID = &collectingTID

	info, err := UpdateAccount(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.AvailableAt = info.AvailableAt
		assert.Equal(t, ret, info)
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
			ID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
			Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
			AvailableAt: &basetypes.Uint32Val{Op: cruder.GT, Value: uint32(time.Now().Unix())},
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
}
