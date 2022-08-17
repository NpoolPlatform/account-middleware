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

func (s *Server) CreateAccount(
	ctx context.Context, in *npool.CreateAccountRequest,
) (
	*npool.CreateAccountResponse, error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "CreateAccount")

	info, err := deposit1.CreateAccount(
		ctx,
		in.GetInfo().GetAppID(),
		in.GetInfo().GetUserID(),
		in.GetInfo().GetCoinTypeID(),
	)
	if err != nil {
		logger.Sugar().Errorf("fail create deposit: %v", err.Error())
		return &npool.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAccountResponse{
		Info: info,
	}, nil
}
