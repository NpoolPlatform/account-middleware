//nolint:nolintlint,dupl
package deposit

import (
	"context"

	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/deposit"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"
)

func (s *Server) UpdateAccount(
	ctx context.Context, in *npool.UpdateAccountRequest,
) (
	*npool.UpdateAccountResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("UpdateAccount", "error", err)
		return &npool.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "UpdateAccount")

	info, err := deposit1.UpdateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create deposit: %v", err.Error())
		return &npool.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) AddAccount(
	ctx context.Context, in *npool.AddAccountRequest,
) (
	*npool.AddAccountResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "AddAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("AddAccount", "error", err)
		return &npool.AddAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "AddAccount")

	info, err := deposit1.AddAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create deposit: %v", err.Error())
		return &npool.AddAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AddAccountResponse{
		Info: info,
	}, nil
}
