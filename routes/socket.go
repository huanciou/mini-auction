package routes

import (
	"mini-auction/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSocketRoutes(router *gin.Engine) {
	server := controllers.SetupSocketIO()

	socketRouters := router.Group("/socket.io")
	{
		socketRouters.GET("/*any", gin.WrapH(server))
		socketRouters.POST("/*any", gin.WrapH(server))
	}
}
