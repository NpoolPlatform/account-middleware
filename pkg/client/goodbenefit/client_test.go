package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"

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
	ID:            uuid.NewString(),
	GoodID:        uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	Address:       uuid.NewString(),
	Backup:        false,
	Active:        true,
	Locked:        false,
	LockedByStr:   accountmgrpb.LockedBy_DefaultLockedBy.String(),
	LockedBy:      accountmgrpb.LockedBy_DefaultLockedBy,
	Blocked:       false,
	TransactionID: uuid1.InvalidUUIDStr,
}

var accReq = npool.AccountReq{
	ID:         &acc.ID,
	GoodID:     &acc.GoodID,
	CoinTypeID: &acc.CoinTypeID,
	Address:    &acc.Address,
	Backup:     &acc.Backup,
	Active:     &acc.Active,
	Locked:     &acc.Locked,
	LockedBy:   &acc.LockedBy,
	Blocked:    &acc.Blocked,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), &accReq)
	if assert.Nil(t, err) {
		acc.CreatedAt = info.CreatedAt
		acc.UpdatedAt = info.UpdatedAt
		acc.AccountID = info.AccountID
		accReq.ID = &info.ID
		accReq.AccountID = &info.AccountID
		assert.Equal(t, info, acc)
	}
}

func updateAccount(t *testing.T) {
	locked := true
	lockedBy := accountmgrpb.LockedBy_Collecting
	blocked := true
	active := false

	acc.Active = active
	acc.Blocked = blocked
	acc.Locked = locked
	acc.LockedBy = lockedBy
	acc.LockedByStr = lockedBy.String()

	accReq.Active = &active
	accReq.Blocked = &blocked
	accReq.Locked = &locked
	accReq.LockedBy = &lockedBy

	info, err := UpdateAccount(context.Background(), &accReq)
	if assert.Nil(t, err) {
		acc.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, acc)
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
