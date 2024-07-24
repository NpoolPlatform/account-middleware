package orderbenefit

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/account-middleware/pkg/db/ent/orderbenefit"
	"github.com/google/uuid"
)

type sqlHandler struct {
	*Handler
	BondAccountID *uuid.UUID
	BondOrderID   *uuid.UUID
	bondVals      map[string]string
	baseVals      map[string]string
	idVals        map[string]string
}

func (h *Handler) newSQLHandler() *sqlHandler {
	return &sqlHandler{
		Handler:  h,
		bondVals: make(map[string]string),
		baseVals: make(map[string]string),
		idVals:   make(map[string]string),
	}
}

//nolint:gocyclo
func (h *sqlHandler) baseKeys() error {
	if h.ID != nil {
		strBytes, err := json.Marshal(*h.ID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldID] = string(strBytes)
	}
	if h.EntID != nil {
		strBytes, err := json.Marshal(*h.EntID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldEntID] = string(strBytes)
	}
	if h.AppID != nil {
		strBytes, err := json.Marshal(*h.AppID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldAppID] = string(strBytes)
	}
	if h.UserID != nil {
		strBytes, err := json.Marshal(*h.UserID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldUserID] = string(strBytes)
	}
	if h.CoinTypeID != nil {
		strBytes, err := json.Marshal(*h.CoinTypeID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldCoinTypeID] = string(strBytes)
	}
	if h.AccountID != nil {
		strBytes, err := json.Marshal(*h.AccountID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldAccountID] = string(strBytes)
		h.BondAccountID = h.AccountID
	}
	if h.OrderID != nil {
		strBytes, err := json.Marshal(*h.OrderID)
		if err != nil {
			return err
		}
		h.baseVals[orderbenefit.FieldOrderID] = string(strBytes)
		h.BondOrderID = h.OrderID
	}

	if h.BondAccountID == nil {
		return fmt.Errorf("please give account id")
	}
	strBytes, err := json.Marshal(*h.BondAccountID)
	if err != nil {
		return err
	}
	h.bondVals[orderbenefit.FieldAccountID] = string(strBytes)

	if h.BondOrderID == nil {
		return fmt.Errorf("please give order id")
	}
	strBytes, err = json.Marshal(*h.BondOrderID)
	if err != nil {
		return err
	}
	h.bondVals[orderbenefit.FieldOrderID] = string(strBytes)
	return nil
}

//nolint:gocognit
func (h *sqlHandler) genCreateSQL() (string, error) {
	err := h.baseKeys()
	if err != nil {
		return "", err
	}
	delete(h.baseVals, orderbenefit.FieldID)

	now := uint32(time.Now().Unix())
	h.baseVals[orderbenefit.FieldCreatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderbenefit.FieldUpdatedAt] = fmt.Sprintf("%v", now)
	h.baseVals[orderbenefit.FieldDeletedAt] = fmt.Sprintf("%v", 0)

	keys := []string{}
	selectVals := []string{}
	bondVals := []string{}

	for k, v := range h.baseVals {
		keys = append(keys, k)
		selectVals = append(selectVals, fmt.Sprintf("%v as %v", v, k))
	}

	for k, v := range h.bondVals {
		bondVals = append(bondVals, fmt.Sprintf("%v=%v", k, v))
	}

	sql := fmt.Sprintf("insert into %v (%v) select * from (select %v) as tmp where not exists (select * from %v where %v and deleted_at=0);",
		orderbenefit.Table,
		strings.Join(keys, ","),
		strings.Join(selectVals, ","),
		orderbenefit.Table,
		strings.Join(bondVals, " AND "),
	)

	return sql, nil
}
