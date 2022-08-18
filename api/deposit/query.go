//nolint:nolintlint,dupl
package deposit

import (
	"context"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/deposit"
)

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	conds := in.GetConds()

	if conds == nil {
		conds = &npool.Conds{}
	}

	infos, err := deposit1.GetAccounts(ctx, conds, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorf("GetAccounts", "err", err)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAccountsResponse{
		Infos: infos,
	}, nil
}
