package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madfelps/go-wallet/schemas"
)

func TransferAmountHandler(ctx *gin.Context) {
	request := CreateTransferRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transfer := schemas.Transfer{
		Amount: request.Amount,
		ID:     request.ID,
		UserID: request.UserID,
	}

	fmt.Println(transfer)

}
