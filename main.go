package main

import (
	"mini-auction/middlewares"
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

	router.Use(middlewares.ErrorHandler())

	routes.RegisterSocketRoutes(router)
	routes.RegisterAuctionRoutes(router)

	router.Run(":3000")
}
