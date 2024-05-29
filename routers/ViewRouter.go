package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewRouterInit(r *gin.Engine) {
	viewRouter := r.Group("/view")
	{
		viewRouter.GET("/upload", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "UpLoadView.html", gin.H{}) })
		viewRouter.GET("/manage", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "ManageView.html", gin.H{}) })
	}
}
