//nolint:dupl
package goodbenefit

import (
	"context"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/goodbenefit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/goodbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(&in.ID),
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
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithConds(in.GetConds()),
		goodbenefit1.WithOffset(in.GetOffset()),
		goodbenefit1.WithLimit(in.GetLimit()),
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
