package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"url_shortner/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Authentication"}

	router.Use(cors.New(config))

	controller := controllers.NewController()

	main := router.Group("api/")
	{
		routers := main.Group("/")
		golyRouter := routers.Group("goly/")
		routers.Use(cors.New(config))
		{
			golyRouter.GET("/", controller.GetAllGolies)
			golyRouter.GET("/:id", controller.GetGoly)
			golyRouter.GET("/url/:uuid", controller.GetGolyUuid)
			golyRouter.DELETE("/:id", controller.DeleteGoly)
			golyRouter.PATCH("/:id", controller.UpdateGoly)
			golyRouter.POST("/", controller.CreateGoly)
			golyRouter.GET("/redirect/:redirect", controller.Redirect)

		}
	}
	return router
}
