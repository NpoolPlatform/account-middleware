package goodbenefit

import (
	"context"

	constant1 "github.com/NpoolPlatform/account-middleware/pkg/const"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/goodbenefit"

	"github.com/google/uuid"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetAccount", "ID", in.GetID(), "error", err)
		return &npool.GetAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := goodbenefit1.GetAccount(ctx, in.GetID())
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

	limit := constant1.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	infos, total, err := goodbenefit1.GetAccounts(ctx, conds, in.GetOffset(), limit)
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

	info, err := goodbenefit1.GetAccountOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAccountOnly", "err", err)
		return &npool.GetAccountOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountOnlyResponse{
		Info: info,
	}, nil
}
