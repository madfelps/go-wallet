package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowWalletHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "show handler",
	})
}
