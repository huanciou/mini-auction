package main

import (
	"mini-auction/models"
	"mini-auction/routes"
	"mini-auction/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	models.RedisInit()
	utils.LoadScripts()
}

func main() {
	router := gin.New()

	routes.RegisterSocketRoutes(router)
	// routes.RegisterAPIRoutes(router)
	routes.RegisterAuctionRoutes(router)

	router.Run(":3000")
}
