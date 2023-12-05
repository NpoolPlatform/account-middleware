//nolint:nolintlint,dupl
package account

import (
	"context"

	account1 "github.com/NpoolPlatform/account-middleware/pkg/mw/account"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/account"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	handler, err := account1.NewHandler(
		ctx,
		account1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	handler, err := account1.NewHandler(
		ctx,
		account1.WithConds(in.GetConds()),
		account1.WithOffset(in.GetOffset()),
		account1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAccounts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
