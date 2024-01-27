package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madfelps/go-wallet/schemas"
)

func ShowWalletsHandler(ctx *gin.Context) {

	wallets := []schemas.Wallet{}
	wallet := schemas.Wallet{}

	sql := `SELECT * FROM wallets`
	rows, err := db.Query(sql)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error selecting wallets on database")
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&wallet.Name, &wallet.Balance)
		if err != nil {
			continue
		}
		wallets = append(wallets, wallet)

	}

	sendSucess(ctx, "show-wallets", wallets)

}
