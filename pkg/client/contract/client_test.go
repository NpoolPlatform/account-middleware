package contract

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"
	accounttypes "github.com/NpoolPlatform/message/npool/basetypes/account/v1"
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
	EntID:                   uuid.NewString(),
	GoodID:                  uuid.NewString(),
	DelegatedStakingID:      uuid.NewString(),
	CoinTypeID:              uuid.NewString(),
	AccountID:               uuid.NewString(),
	Address:                 uuid.NewString(),
	Backup:                  false,
	Active:                  true,
	Locked:                  false,
	LockedByStr:             basetypes.AccountLockedBy_DefaultLockedBy.String(),
	Blocked:                 false,
	ContractOperatorType:    accounttypes.ContractOperatorType_ContractOwner,
	ContractOperatorTypeStr: accounttypes.ContractOperatorType_ContractOwner.String(),
}

var req = npool.AccountReq{
	EntID:                &ret.EntID,
	GoodID:               &ret.GoodID,
	DelegatedStakingID:   &ret.DelegatedStakingID,
	ContractOperatorType: &ret.ContractOperatorType,
	CoinTypeID:           &ret.CoinTypeID,
	AccountID:            &ret.AccountID,
	Address:              &ret.Address,
	Backup:               &ret.Backup,
	Active:               &ret.Active,
	Locked:               &ret.Locked,
	Blocked:              &ret.Blocked,
}

func createAccount(t *testing.T) {
	_, err := CreateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		info, err := GetAccount(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
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

	req.ID = &ret.ID
	req.Active = &active
	req.Blocked = &blocked
	req.Locked = &locked
	req.LockedBy = &lockedBy

	_, err := UpdateAccount(context.Background(), &req)
	if assert.Nil(t, err) {
		info, err := GetAccount(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getAccount(t *testing.T) {
	info, err := GetAccount(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAccounts(t *testing.T) {
	infos, total, err := GetAccounts(
		context.Background(),
		&npool.Conds{
			EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAccount", createAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("deleteAccount", deleteAccount)
}
