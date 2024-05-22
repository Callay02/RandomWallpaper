/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-21 12:39:19
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-21 13:04:46
 * @FilePath: \randomWallpaper\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
