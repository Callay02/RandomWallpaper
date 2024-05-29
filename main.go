package main

import (
	"RandomWallpaper/config"
	"RandomWallpaper/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	routers.ApiRouterInit(r)
	routers.ViewRouterInit(r)
	r.Run(":" + config.Cfg.Server.Port)
}
