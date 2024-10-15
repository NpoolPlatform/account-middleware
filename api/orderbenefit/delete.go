package orderbenefit

import (
	"context"

	orderbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/orderbenefit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/orderbenefit"
)

func (s *Server) DeleteAccount(ctx context.Context, in *npool.DeleteAccountRequest) (*npool.DeleteAccountResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAccount",
			"In", in,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := orderbenefit1.NewHandler(
		ctx,
		orderbenefit1.WithID(req.ID, false),
		orderbenefit1.WithEntID(req.EntID, false),
		orderbenefit1.WithAppID(req.AppID, true),
		orderbenefit1.WithUserID(req.UserID, true),
		orderbenefit1.WithOrderID(req.OrderID, true),
		orderbenefit1.WithAccountID(req.AccountID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}

	info, err := handler.GetAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}
	if info == nil {
		return &npool.DeleteAccountResponse{}, nil
	}

	err = handler.DeleteAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.DeleteAccountResponse{
		Info: info,
	}, nil
}

//nolint:dupl
func (s *Server) DeleteAccounts(ctx context.Context, in *npool.DeleteAccountsRequest) (*npool.DeleteAccountsResponse, error) {
	handler, err := orderbenefit1.NewMultiDeleteHandler(
		ctx,
		in.GetInfos(),
		true,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	err = handler.DeleteAccounts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccounts",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.DeleteAccountsResponse{}, nil
}
