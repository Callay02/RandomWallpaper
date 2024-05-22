package services

import (
	"RandomWallpaper/config"
	"RandomWallpaper/models"
	"RandomWallpaper/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiService struct {
}

func (api ApiService) Random(ctx *gin.Context) {
	s := ctx.Query("source")
	source := strings.Split(s, ",")[0]
	t := ctx.Query("type")
	imgType := strings.Split(t, ",")[0]
	image := models.Image{}

	if imgType == "" {
		if source == "" {
			utils.DB.Raw("select * from images where state = 1 order by rand() limit 1").Scan(&image)
			if image.Source == "oss" {
				ctx.Redirect(http.StatusMovedPermanently, image.Url)
			} else {
				ctx.File(image.Url)
			}
		} else if source == "local" {
			utils.DB.Raw("select * from images where state = 1 and source = 'local' order by rand() limit 1").Scan(&image)
			ctx.File(image.Url)
		} else if source == "oss" {
			utils.DB.Raw("select * from images where state = 1 and source = 'oss' order by rand() limit 1").Scan(&image)
			ctx.Redirect(http.StatusMovedPermanently, image.Url)
		}
	} else {
		if source == "" {
			utils.DB.Raw("select * from images where state = 1 and type = ? order by rand() limit 1", imgType).Scan(&image)
			if image.Source == "oss" {
				ctx.Redirect(http.StatusMovedPermanently, image.Url)
			} else {
				ctx.File(image.Url)
			}
		} else if source == "local" {
			utils.DB.Raw("select * from images where state = 1 and source = 'local' and type = ? order by rand() limit 1", imgType).Scan(&image)
			ctx.File(image.Url)
		} else if source == "oss" {
			utils.DB.Raw("select * from images where state = 1 and source = 'oss' and type = ? order by rand() limit 1", imgType).Scan(&image)
			ctx.Redirect(http.StatusMovedPermanently, image.Url)
		}
	}
}

func (api ApiService) UpLoad(ctx *gin.Context) {
	var image models.Image

	image.Info = ctx.PostForm("info")
	image.Type = ctx.PostForm("type")
	image.Source = ctx.PostForm("source")
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
		//fmt.Println(image.Url)
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
			ctx.JSON(http.StatusOK, new(models.Result).Error("上传失败"))
			return
		}
		if err := utils.DB.Create(image).Error; err != nil {
			ctx.JSON(http.StatusOK, new(models.Result).Error("数据库存入失败"))
			return
		}
		ctx.JSON(http.StatusOK, new(models.Result).OK("上传成功"))
	}()
}
