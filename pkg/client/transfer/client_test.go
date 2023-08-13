package transfer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"
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

var ret = npool.Transfer{
	ID:           uuid.NewString(),
	AppID:        uuid.NewString(),
	UserID:       uuid.NewString(),
	TargetUserID: uuid.NewString(),
}

var req = npool.TransferReq{
	ID:           &ret.ID,
	AppID:        &ret.AppID,
	UserID:       &ret.UserID,
	TargetUserID: &ret.TargetUserID,
}

func createTransfer(t *testing.T) {
	info, err := CreateTransfer(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getTransfer(t *testing.T) {
	info, err := GetTransfer(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTransfers(t *testing.T) {
	infos, total, err := GetTransfers(
		context.Background(),
		&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			TargetUserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.TargetUserID},
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

func deleteTransfer(t *testing.T) {
	info, err := DeleteTransfer(context.Background(), &npool.TransferReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = DeleteTransfer(context.Background(), &npool.TransferReq{
		ID: &ret.ID,
	})
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createTransfer", createTransfer)
	t.Run("getTransfer", getTransfer)
	t.Run("getTransfers", getTransfers)
	t.Run("deleteTransfer", deleteTransfer)
}
