package platform

import (
	"context"

	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	platform1 "github.com/NpoolPlatform/account-middleware/pkg/platform"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

	"github.com/google/uuid"
)

func (s *Server) UpdateAccount(ctx context.Context, in *npool.UpdateAccountRequest) (*npool.UpdateAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAccount", "ID", in.GetInfo().GetID(), "err", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "platform", "platform", "UpdateAccount")

	info, err := platform1.UpdateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAccount", "err", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: info,
	}, nil
}
