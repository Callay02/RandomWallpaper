package routers

import (
	"RandomWallpaper/config"
	"RandomWallpaper/middlewares"
	"RandomWallpaper/services"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiRouter := r.Group("/api")
	validate := middlewares.Validate{}
	if config.Cfg.Server.Token {
		apiRouter.Use(validate.Token)
	}
	{
		apiService := services.ApiService{}
		apiRouter.GET("/random", apiService.Random)
		apiRouter.POST("/upload", apiService.UpLoad)
	}
}
