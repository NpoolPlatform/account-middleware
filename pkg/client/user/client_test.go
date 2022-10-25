package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

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

var acc = &npool.Account{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	UserID:     uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	Address:    uuid.NewString(),
	Active:     true,
	UsedFor:    accountmgrpb.AccountUsedFor_UserWithdraw,
	UsedForStr: accountmgrpb.AccountUsedFor_UserWithdraw.String(),
	Labels:     []string{uuid.NewString(), uuid.NewString()},
}

var accReq = &npool.AccountReq{
	ID:         &acc.ID,
	AppID:      &acc.AppID,
	UserID:     &acc.UserID,
	CoinTypeID: &acc.CoinTypeID,
	Address:    &acc.Address,
	UsedFor:    &acc.UsedFor,
	Labels:     acc.Labels,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), accReq)
	if assert.Nil(t, err) {
		acc.ID = info.ID
		acc.CreatedAt = info.CreatedAt
		acc.LabelsStr = info.LabelsStr
		acc.AccountID = info.AccountID
		assert.Equal(t, acc, info)
	}
}

func updateAccount(t *testing.T) {
	active := false
	labels := []string{uuid.NewString(), uuid.NewString()}
	blocked := true

	acc.Active = active
	acc.Labels = labels
	acc.Blocked = blocked

	accReq.Active = &active
	accReq.Labels = labels
	accReq.Blocked = &blocked

	info, err := UpdateAccount(context.Background(), accReq)
	if assert.Nil(t, err) {
		acc.LabelsStr = info.LabelsStr
		assert.Equal(t, acc, info)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction { //nolint
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAccount", createAccount)
	t.Run("updateAccount", updateAccount)
}
