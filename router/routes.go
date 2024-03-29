package router

import (
	"github.com/gin-gonic/gin"
	"github.com/madfelps/go-wallet/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallet", handler.ShowWalletHandler)
		v1.POST("/wallet", handler.CreateWalletHandler)
		v1.GET("/wallets", handler.ShowWalletsHandler)
		v1.POST("/transfer", handler.TransferAmountHandler)
	}
}
