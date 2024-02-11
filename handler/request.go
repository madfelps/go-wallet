package handler

import (
	"fmt"

	"github.com/google/uuid"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errParamMismatch(name string) error {
	return fmt.Errorf("param: %s value doesn't make sense", name)
}

type CreateWalletRequest struct {
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type CreateTransferRequest struct {
	Amount int       `json:"amount"`
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

func (r *CreateWalletRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Balance == "" {
		return errParamIsRequired("balance", "string")
	}
	return nil
}

func (r *CreateTransferRequest) Validate() error {
	if r.Amount < 0 {
		return errParamMismatch("amount")
	}
	return nil
}
