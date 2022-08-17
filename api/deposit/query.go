//nolint:nolintlint,dupl
package deposit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/deposit"
)

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	infos, err := deposit1.GetAccounts(ctx, in.GetConds())
	if err != nil {
		return nil, err
	}
	return &npool.GetAccountsResponse{
		Infos: infos,
	}, nil
}