package user

import (
	"context"

	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	user1 "github.com/NpoolPlatform/account-middleware/pkg/user"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

	"github.com/google/uuid"
)

func (s *Server) DeleteAccount(ctx context.Context, in *npool.DeleteAccountRequest) (*npool.DeleteAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteAccount", "ID", in.GetID(), "err", err)
		return &npool.DeleteAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "user", "user", "DeleteAccount")

	info, err := user1.DeleteAccount(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteAccount", "err", err)
		return &npool.DeleteAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAccountResponse{
		Info: info,
	}, nil
}
