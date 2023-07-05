package deposit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	AccountID:     uuid.NewString(),
	Address:       uuid.NewString(),
	Active:        true,
	Locked:        false,
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
	Blocked:       false,
	CollectingTID: uuid.UUID{}.String(),
	Incoming:      decimal.NewFromInt(0).String(),
	Outcoming:     decimal.NewFromInt(0).String(),
}

var req = npool.AccountReq{
	ID:         &ret.ID,
	AppID:      &ret.AppID,
	UserID:     &ret.UserID,
	CoinTypeID: &ret.CoinTypeID,
	AccountID:  &ret.AccountID,
	Address:    &ret.Address,
	Active:     &ret.Active,
	Locked:     &ret.Locked,
	Blocked:    &ret.Blocked,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ScannableAt = info.ScannableAt
		assert.Equal(t, info, ret)
	}
}

func updateAccount(t *testing.T) {
	collectingTID := uuid.NewString()
	incoming := "1.2"
	outcoming := "1.1"

	ret.CollectingTID = collectingTID
	ret.Incoming = incoming
	ret.Outcoming = outcoming
	ret.Locked = true
	ret.LockedBy = basetypes.AccountLockedBy_Payment
	ret.LockedByStr = basetypes.AccountLockedBy_Payment.String()

	req.CollectingTID = &collectingTID
	req.Incoming = &incoming
	req.Outcoming = &outcoming
	req.LockedBy = &ret.LockedBy

	info, err := UpdateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	ret.Locked = false
	req.Locked = &ret.Locked
	req.Incoming = nil
	req.Outcoming = nil

	info, err = UpdateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.NotEqual(t, info.ScannableAt, ret.ScannableAt)
		ret.ScannableAt = info.ScannableAt
		assert.Equal(t, info, ret)
	}
}

func getAccount(t *testing.T) {
	info, err := GetAccount(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getAccounts(t *testing.T) {
	infos, total, err := GetAccounts(
		context.Background(),
		&npool.Conds{
			ID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Locked},
			Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
			ScannableAt: &basetypes.Uint32Val{Op: cruder.GT, Value: uint32(time.Now().Unix())},
		},
		0,
		int32(2),
	)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, uint32(1)) {
			assert.Equal(t, infos[0], ret)
		}
	}
}

func deleteAccount(t *testing.T) {
	info, err := DeleteAccount(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	info, err = DeleteAccount(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Nil(t, info)
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
