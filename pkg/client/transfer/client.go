package transfer

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"

	servicename "github.com/NpoolPlatform/account-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
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

func CreateTransfer(ctx context.Context, in *npool.TransferReq) (*npool.Transfer, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateTransfer(ctx, &npool.CreateTransferRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Transfer), nil
}

func GetTransfer(ctx context.Context, id string) (*npool.Transfer, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTransfer(ctx, &npool.GetTransferRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Transfer), nil
}

func GetTransfers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Transfer, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTransfers(ctx, &npool.GetTransfersRequest{
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
	return infos.([]*npool.Transfer), total, nil
}

func GetTransferOnly(ctx context.Context, conds *npool.Conds) (*npool.Transfer, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetTransfers(ctx, &npool.GetTransfersRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Transfer)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Transfer)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Transfer)[0], nil
}

func DeleteTransfer(ctx context.Context, in *npool.TransferReq) (*npool.Transfer, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteTransfer(ctx, &npool.DeleteTransferRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Transfer), nil
}

func ExistTransfer(ctx context.Context, id string) (bool, error) {
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistTransfer(ctx, &npool.ExistTransferRequest{
			EntID: id,
		})
		if err != nil {
			return false, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}

func ExistTransferConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistTransferConds(ctx, &npool.ExistTransferCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return false, err
		}

		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}
