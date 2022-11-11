//nolint:nolintlint,dupl
package account

import (
	"context"

	constant1 "github.com/NpoolPlatform/account-middleware/pkg/const"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	accmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	accmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/account"
)

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

	conds := in.Conds
	if conds == nil {
		conds = &accmgrpb.Conds{}
	}

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultRowLimit
	}

	span = commontracer.TraceInvoker(span, "deposit", "deposit", "GetAccounts")

	infos, total, err := accmgrcli.GetAccounts(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetAccounts", "err", err)
		return &npool.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAccountsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
