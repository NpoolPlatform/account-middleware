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

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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
	Address:       uuid.NewString(),
	Active:        true,
	Locked:        false,
	LockedByStr:   basetypes.AccountLockedBy_DefaultLockedBy.String(),
	LockedBy:      basetypes.AccountLockedBy_DefaultLockedBy,
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
		acc.AccountID = info.AccountID
		assert.Equal(t, info, acc)
	}
}

func updateAccount(t *testing.T) {
	collectingTID := uuid.NewString()
	incoming := "1.200000000000000000"
	outcoming := "1.100000000000000000"
	scannableAt := uint32(time.Now().Unix() + 100000)

	accReq.CollectingTID = &collectingTID
	accReq.Incoming = &incoming
	accReq.Outcoming = &outcoming
	accReq.ScannableAt = &scannableAt

	acc.CollectingTID = collectingTID
	acc.Incoming = incoming
	acc.Outcoming = outcoming
	acc.ScannableAt = scannableAt

	info, err := UpdateAccount(context.Background(), &accReq)
	if assert.Nil(t, err) {
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
