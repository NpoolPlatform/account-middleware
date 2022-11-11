package account

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	accmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/account"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetAccounts(ctx context.Context, conds *accmgrpb.Conds, offset, limit int32) ([]*accmgrpb.Account, uint32, error) {
	total := uint32(0)

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAccounts(ctx, &npool.GetAccountsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*accmgrpb.Account), total, nil
}
