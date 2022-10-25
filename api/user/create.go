//nolint:dupl
package user

import (
	"context"
	"fmt"

	commontracer "github.com/NpoolPlatform/account-middleware/pkg/tracer"

	accountmgrcli "github.com/NpoolPlatform/account-manager/pkg/client/account"
	constant "github.com/NpoolPlatform/account-middleware/pkg/message/const"
	user1 "github.com/NpoolPlatform/account-middleware/pkg/user"

	commonpb "github.com/NpoolPlatform/message/npool"
	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/user"

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

	switch in.GetInfo().GetUsedFor() {
	case accountmgrpb.AccountUsedFor_UserWithdraw:
	case accountmgrpb.AccountUsedFor_UserDirectBenefit:
	default:
		logger.Sugar().Errorw("CreateAccount", "UsedFor", in.GetInfo().GetUsedFor())
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "AppID", in.GetInfo().GetAppID(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}
	if _, err := uuid.Parse(in.GetInfo().GetUserID()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "UserID", in.GetInfo().GetUserID(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}
	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateAccount", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}
	if in.GetInfo().GetAddress() == "" {
		logger.Sugar().Errorw("CreateAccount", "Address", in.GetInfo().GetAddress(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Address is invalid: %v", err))
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
		logger.Sugar().Errorw("CreateAccount", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "Address", in.GetInfo().GetAddress(), "error", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Address is invalid: %v", err))
	}
	if exist {
		logger.Sugar().Errorw("CreateAccount", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "Address", in.GetInfo().GetAddress(), "exist", exist)
		return &npool.CreateAccountResponse{}, status.Error(codes.AlreadyExists, "Address already exists")
	}

	span = commontracer.TraceInvoker(span, "user", "user", "CreateAccount")

	info, err := user1.CreateAccount(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "err", err)
		return &npool.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAccountResponse{
		Info: info,
	}, nil
}
