package payment

import (
	"context"

	constant1 "github.com/NpoolPlatform/account-middleware/pkg/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/payment"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	var err error

	info, err := payment1.GetAccount(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "err", err)
		return &npool.GetAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	var err error

	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	limit := constant1.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := payment1.GetAccounts(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetAccounts", "err", err)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAccountOnly(ctx context.Context, in *npool.GetAccountOnlyRequest) (*npool.GetAccountOnlyResponse, error) {
	var err error

	info, err := payment1.GetAccountOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAccountOnly", "err", err)
		return &npool.GetAccountOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountOnlyResponse{
		Info: info,
	}, nil
}
