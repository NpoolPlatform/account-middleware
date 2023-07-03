package account

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	servicename "github.com/NpoolPlatform/account-middleware/pkg/servicename"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetAccount(ctx context.Context, id string) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAccount(ctx, &npool.GetAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Account), nil
}

func GetAccounts(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Account, uint32, error) {
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
	return infos.([]*npool.Account), total, nil
}

func GetManyAccounts(ctx context.Context, ids []string) ([]*npool.Account, uint32, error) {
	total := uint32(0)

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAccounts(ctx, &npool.GetAccountsRequest{
			Conds: &npool.Conds{
				IDs: &basetypes.StringSliceVal{
					Op:    cruder.IN,
					Value: ids,
				},
			},
			Offset: 0,
			Limit:  int32(len(ids)),
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
	return infos.([]*npool.Account), total, nil
}
