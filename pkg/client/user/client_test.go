package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"
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
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	UserID:     uuid.NewString(),
	AccountID:  uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	Address:    uuid.NewString(),
	Active:     true,
	UsedFor:    basetypes.AccountUsedFor_UserWithdraw,
	UsedForStr: basetypes.AccountUsedFor_UserWithdraw.String(),
	Labels:     []string{uuid.NewString(), uuid.NewString()},
	Memo:       uuid.NewString(),
}

var retReq = &npool.AccountReq{
	ID:         &ret.ID,
	AppID:      &ret.AppID,
	UserID:     &ret.UserID,
	AccountID:  &ret.AccountID,
	CoinTypeID: &ret.CoinTypeID,
	Address:    &ret.Address,
	UsedFor:    &ret.UsedFor,
	Labels:     ret.Labels,
	Memo:       &ret.Memo,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.LabelsStr = info.LabelsStr
		ret.UpdatedAt = info.UpdatedAt
		ret.DeletedAt = info.DeletedAt
		assert.Equal(t, ret, info)
	}
}

func updateAccount(t *testing.T) {
	active := false
	labels := []string{uuid.NewString(), uuid.NewString()}
	blocked := true
	memo := uuid.NewString()

	ret.Active = active
	ret.Labels = labels
	ret.Blocked = blocked
	ret.Memo = memo

	retReq.Active = &active
	retReq.Labels = labels
	retReq.Blocked = &blocked
	retReq.Memo = &memo

	info, err := UpdateAccount(context.Background(), retReq)
	if assert.Nil(t, err) {
		ret.LabelsStr = info.LabelsStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.DeletedAt = info.DeletedAt
		assert.Equal(t, ret, info)
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
			ID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AccountID},
			Address:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.Address},
			Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Active},
			Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Blocked},
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
	req := &npool.AccountReq{
		ID: &ret.ID,
	}
	info, err := DeleteAccount(context.Background(), req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
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
