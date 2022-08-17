//nolint:nolintlint,dupl
package deposit

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
