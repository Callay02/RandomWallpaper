package middlewares

import (
	"RandomWallpaper/models"
	"RandomWallpaper/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Validate struct{}

func (Validate) Token(ctx *gin.Context) {
	u := models.User{}
	var token string
	switch ctx.Request.Method {
	case "GET":
		token = strings.Split(ctx.Query("token"), ",")[0]
	case "POST":
		token = ctx.PostForm("token")
	}
	if err := utils.DB.Where("token = ? and state = ?", token, 1).First(&u).Error; err != nil {
		ctx.AbortWithStatusJSON(200, new(utils.Result).Error("token错误"))
	}
	fmt.Println(u)
	ctx.Set("uid", u.Id)
}

func (Validate) Role(ctx *gin.Context) {
	u := models.User{}
	r := models.Role{}
	var token string
	switch ctx.Request.Method {
	case "GET":
		token = strings.Split(ctx.Query("token"), ",")[0]
	case "POST":
		token = ctx.PostForm("token")
	}
	if err := utils.DB.Where("token = ? and state = ?", token, 1).First(&u).Error; err != nil {
		ctx.AbortWithStatusJSON(200, new(utils.Result).Error("token错误"))
	}
	if err := utils.DB.Where("id = ? and state = ?", u.Role, 1).First(&r).Error; err != nil {
		ctx.AbortWithStatusJSON(200, new(utils.Result).Error("该用户没有相应权限"))
	}
	ctx.Set("uid", u.Id)
	ctx.Set("role", r.Name)

}
