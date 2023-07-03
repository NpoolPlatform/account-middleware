package payment

import (
	"context"
	"time"

	accountcrud "github.com/NpoolPlatform/account-manager/pkg/crud/account"
	paymentcrud "github.com/NpoolPlatform/account-manager/pkg/crud/payment"
	"github.com/NpoolPlatform/account-manager/pkg/db"
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	entaccount "github.com/NpoolPlatform/account-manager/pkg/db/ent/account"
	entpayment "github.com/NpoolPlatform/account-manager/pkg/db/ent/payment"

	accountmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	paymentmgrpb "github.com/NpoolPlatform/message/npool/account/mgr/v1/payment"
	npool "github.com/NpoolPlatform/message/npool/account/mw/v1/payment"

	"github.com/google/uuid"
)

func UpdateAccount(ctx context.Context, in *npool.AccountReq) (info *npool.Account, err error) {
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		payment, err := tx.Payment.
			Query().
			Where(
				entpayment.ID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.ID(payment.AccountID),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if account.Locked && !in.GetLocked() {
			const coolDown = uint32(60 * 60)
			availableAt := uint32(time.Now().Unix()) + coolDown
			in.AvailableAt = &availableAt
		}

		if _, err = paymentcrud.UpdateSet(payment, &paymentmgrpb.AccountReq{
			AccountID:     in.AccountID,
			CollectingTID: in.CollectingTID,
			AvailableAt:   in.AvailableAt,
		}).Save(ctx); err != nil {
			return err
		}

		if _, err := accountcrud.UpdateSet(account, &accountmgrpb.AccountReq{
			Active:   in.Active,
			Locked:   in.Locked,
			LockedBy: in.LockedBy,
			Blocked:  in.Blocked,
		}).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetAccount(ctx, in.GetID())
}
