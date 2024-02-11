package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madfelps/go-wallet/schemas"
)

func ShowWalletsHandler(ctx *gin.Context) {

	wallets := []schemas.Wallet{}
	wallet := schemas.Wallet{}

	query := `SELECT * FROM wallets`
	rows, err := db.Query(query)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error selecting wallets on database")
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&wallet.Id, &wallet.Name, &wallet.Balance)
		if err != nil {
			continue
		}
		fmt.Printf("Name: %s", wallet.Name)
		fmt.Printf("Balance: %s", wallet.Balance)
		fmt.Printf("Id: %d", wallet.Id)
		if err != nil {
			continue
		}
		wallets = append(wallets, wallet)

	}

	sendSucess(ctx, "show-wallets", wallets)

}
