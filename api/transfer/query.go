//nolint:nolintlint,dupl
package transfer

import (
	"context"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/transfer"

	transfermgrcli "github.com/NpoolPlatform/account-manager/pkg/client/transfer"

	"github.com/google/uuid"
)

func (s *Server) ExistTransferConds(ctx context.Context, in *npool.ExistTransferCondsRequest) (*npool.ExistTransferCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTransfer")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.Conds == nil {
		logger.Sugar().Errorw("ExistTransferConds", "Conds", in.GetConds())
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetConds().GetUserID().GetValue()); err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "UserID", in.GetConds().GetUserID().GetValue(), "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetConds().GetTargetUserID().GetValue()); err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "TargetUserID", in.GetConds().GetTargetUserID().GetValue(), "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "transfer", "transfer", "ExistTransferConds")

	exist, err := transfermgrcli.ExistTransferConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "err", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferCondsResponse{
		Info: exist,
	}, nil
}
