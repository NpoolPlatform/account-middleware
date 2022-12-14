//nolint:nolintlint,dupl
package deposit

import (
	"context"

	constant1 "github.com/NpoolPlatform/account-middleware/pkg/const"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

	deposit1 "github.com/NpoolPlatform/account-middleware/pkg/deposit"
)

func (s *Server) GetAccount(ctx context.Context, in *npool.GetAccountRequest) (*npool.GetAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetAccount", "error", err)
		return &npool.GetAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "GetAccount")

	info, err := deposit1.GetAccount(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "err", err)
		return &npool.GetAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *npool.GetAccountsRequest) (*npool.GetAccountsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultRowLimit
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "GetAccounts")

	infos, total, err := deposit1.GetAccounts(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetAccounts", "err", err)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
