package routers

import (
	"RandomWallpaper/services"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiService := services.ApiService{}
		apiRouter.GET("/random", func(ctx *gin.Context) {
			go apiService.Random(ctx)
		})
		apiRouter.POST("/upload", func(ctx *gin.Context) {
			go apiService.UpLoad(ctx)
		})
	}
}
