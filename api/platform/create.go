package platform

import (
	"context"

	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	commonpb "github.com/NpoolPlatform/message/npool"

	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	platform1 "github.com/NpoolPlatform/account-middleware/pkg/platform"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/platform"

	"github.com/google/uuid"
)

func (s *Server) CreateAccount(ctx context.Context, in *npool.CreateAccountRequest) (*npool.CreateAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAccount")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			logger.Sugar().Errorw("CreateAccount", "ID", in.GetInfo().GetID(), "err", err)
			return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().GetAddress() == "" {
		logger.Sugar().Errorw("CreateAccount", "Address", in.GetInfo().GetAddress(), "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	switch in.GetInfo().GetUsedFor() {
	case accountmgrpb.AccountUsedFor_UserBenefitHot:
	case accountmgrpb.AccountUsedFor_UserBenefitCold:
	case accountmgrpb.AccountUsedFor_PlatformBenefitCold:
	case accountmgrpb.AccountUsedFor_GasProvider:
	case accountmgrpb.AccountUsedFor_PaymentCollector:
	default:
		logger.Sugar().Errorw("CreateAccount", "UsedFor", in.GetInfo().GetUsedFor(), "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := accountmgrcli.ExistAccountConds(ctx, &accountmgrpb.Conds{
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetCoinTypeID(),
		},
		Address: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetAddress(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "Address", in.GetInfo().GetAddress(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("validate", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "Address", in.GetInfo().GetAddress(), "exist", exist)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "platform", "platform", "CreateAccount")

	info, err := platform1.CreateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAccountResponse{
		Info: info,
	}, nil
}
