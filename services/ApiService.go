package services

import (
	"RandomWallpaper/config"
	"RandomWallpaper/models"
	"RandomWallpaper/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiService struct {
}

func (ApiService) Random(ctx *gin.Context) {
	var sourceSQL string
	var typeSQL string
	image := models.Image{}

	if ctx.Query("source") != "" {
		sourceSQL = fmt.Sprintf("and source = '%v'", strings.Split(ctx.Query("source"), ",")[0])
	}
	if ctx.Query("type") != "" {
		typeSQL = fmt.Sprintf("and type = '%v'", strings.Split(ctx.Query("type"), ",")[0])
	}
	sql := fmt.Sprintf("select * from images where state = 1 %v %v order by rand() limit 1", sourceSQL, typeSQL)

	utils.DB.Raw(sql).Scan(&image)

	switch image.Source {
	case "local":
		ctx.File(image.Url)
	case "oss":
		ctx.Redirect(http.StatusMovedPermanently, image.Url)
	}
}

func (ApiService) UpLoad(ctx *gin.Context) {
	var image models.Image

	if uid, ok := ctx.Get("uid"); ok {
		if v, ok := uid.(int); ok {
			image.User = v
		}
	}
	image.Source = ctx.PostForm("source")
	image.Info = ctx.PostForm("info")
	image.Type = ctx.PostForm("type")
	image.UpdateTime = strconv.FormatInt(time.Now().Unix(), 10)
	image.State = 1

	fh, _ := ctx.FormFile("img")

	if image.Source == "oss" {
		//转为文件流
		f, err := fh.Open()
		if err != nil {
			image.State = 0
			return
		}
		//执行上传阿里云OSS
		oss := utils.OSS{}
		if err := oss.UpLoad(&image, f); err != nil {
			image.State = 0
			return
		}
	} else {
		//本地
		image.Source = "local"
		image.Url = config.Cfg.Server.ImgSavePath + "/" + image.Type + "/" + image.UpdateTime + "-" + image.Info
		if err := ctx.SaveUploadedFile(fh, image.Url); err != nil {
			image.State = 0
			return
		}
	}

	defer func() {
		if image.State == 0 {
			ctx.JSON(http.StatusOK, new(utils.Result).Error("上传失败"))
			return
		}
		if err := utils.DB.Create(image).Error; err != nil {
			ctx.JSON(http.StatusOK, new(utils.Result).Error("数据库存入失败"))
			return
		}
		ctx.JSON(http.StatusOK, new(utils.Result).OK("上传成功"))
	}()
}
