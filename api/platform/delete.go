package platform

import (
	"context"

	platform1 "github.com/NpoolPlatform/account-middleware/pkg/mw/platform"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

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
	handler, err := platform1.NewHandler(
		ctx,
		platform1.WithID(req.ID),
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
