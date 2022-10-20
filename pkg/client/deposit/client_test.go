package deposit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

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
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	CoinTypeID:    uuid.NewString(),
	AccountID:     uuid.NewString(),
	Address:       uuid.NewString(),
	Active:        true,
	Locked:        false,
	LockedByStr:   accountmgrpb.LockedBy_DefaultLockedBy.String(),
	LockedBy:      accountmgrpb.LockedBy_DefaultLockedBy,
	Blocked:       false,
	CollectingTID: uuid.UUID{}.String(),
	Incoming:      "0.000000000000000000",
	Outcoming:     "0.000000000000000000",
}

var accReq = npool.AccountReq{
	ID:         &acc.ID,
	AppID:      &acc.AppID,
	UserID:     &acc.UserID,
	CoinTypeID: &acc.CoinTypeID,
	AccountID:  &acc.AccountID,
	Address:    &acc.Address,
	Active:     &acc.Active,
	Locked:     &acc.Locked,
	LockedBy:   &acc.LockedBy,
	Blocked:    &acc.Blocked,
}

func createAccount(t *testing.T) {
	info, err := CreateAccount(context.Background(), &accReq)
	if assert.Nil(t, err) {
		acc.CreatedAt = info.CreatedAt
		acc.ScannableAt = info.ScannableAt
		assert.Equal(t, info, acc)
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
}
