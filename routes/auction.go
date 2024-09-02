package routes

import (
	"mini-auction/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuctionRoutes(router *gin.Engine) {
	auctionRouters := router.Group("/auction")
	{
		auctionRouters.GET("/", controllers.Auction)
	}
}
