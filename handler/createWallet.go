package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWalletHandler(ctx *gin.Context) {
	request := CreateWalletRequest{}

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.BindJSON(&request)
}
