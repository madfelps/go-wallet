package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateWalletRequest struct {
	Name    string `json:"name"`
	Balance string `json:"balance"`
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
