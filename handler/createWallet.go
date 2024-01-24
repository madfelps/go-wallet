package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWalletHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok!",
	})
}
