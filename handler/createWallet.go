package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madfelps/go-wallet/schemas"
)

func CreateWalletHandler(ctx *gin.Context) {
	request := CreateWalletRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	wallet := schemas.Wallet{
		Name:    request.Name,
		Balance: request.Balance,
	}

	//popular o banco
	fmt.Println(wallet.Name, wallet.Balance)
	sql := `INSERT INTO wallets (name, balance) VALUES($1, $2)`
	result, err := db.Exec(sql, wallet.Name, wallet.Balance)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error creating wallet on database")
		return
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows Affected:", rowsAffected)

	sendSucess(ctx, "create-wallet", wallet)
}
