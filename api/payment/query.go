package payment

import (
	"context"

	payment1 "github.com/NpoolPlatform/account-middleware/pkg/mw/payment"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithConds(in.GetConds()),
		payment1.WithOffset(in.GetOffset()),
		payment1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetAccounts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAccountOnly(ctx context.Context, in *npool.GetAccountOnlyRequest) (*npool.GetAccountOnlyResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccountOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetAccountOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAccountOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetAccountOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAccountOnlyResponse{
		Info: info,
	}, nil
}
