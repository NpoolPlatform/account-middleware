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
