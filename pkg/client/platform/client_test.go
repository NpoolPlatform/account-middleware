package platform

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

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
	ID:          uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
	UsedFor:     accountmgrpb.AccountUsedFor_UserBenefitHot,
	UsedForStr:  accountmgrpb.AccountUsedFor_UserBenefitHot.String(),
	Address:     uuid.NewString(),
	Backup:      false,
	Active:      true,
	Locked:      false,
	LockedBy:    accountmgrpb.LockedBy_DefaultLockedBy,
	LockedByStr: accountmgrpb.LockedBy_DefaultLockedBy.String(),
	Blocked:     false,
}

var accReq = &npool.AccountReq{
	ID:         &acc.ID,
	CoinTypeID: &acc.CoinTypeID,
	UsedFor:    &acc.UsedFor,
	Address:    &acc.Address,
	LockedBy:   &acc.LockedBy,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), accReq)
	if assert.Nil(t, err) {
		acc.CreatedAt = info.CreatedAt
		acc.AccountID = info.AccountID
		assert.Equal(t, acc, info)
	}
}

func updateAccount(t *testing.T) {
	backup := true
	locked := true
	lockedBy := accountmgrpb.LockedBy_Collecting
	blocked := true

	acc.Backup = backup
	acc.Locked = locked
	acc.LockedBy = lockedBy
	acc.LockedByStr = lockedBy.String()
	acc.Blocked = blocked

	accReq.Backup = &backup
	accReq.Locked = &locked
	accReq.LockedBy = &lockedBy
	accReq.Blocked = &blocked

	info, err := UpdateAccount(context.Background(), accReq)
	if assert.Nil(t, err) {
		acc.CreatedAt = info.CreatedAt
		assert.Equal(t, acc, info)
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
}
