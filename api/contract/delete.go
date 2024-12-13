package contract

import (
	"context"

	goodbenefit1 "github.com/NpoolPlatform/account-middleware/pkg/mw/contract"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/contract"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAccount(ctx context.Context, in *npool.DeleteAccountRequest) (*npool.DeleteAccountResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := goodbenefit1.NewHandler(
		ctx,
		goodbenefit1.WithID(req.ID, false),
		goodbenefit1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAccount",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAccountResponse{
		Info: info,
	}, nil
}
