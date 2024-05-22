package utils

import (
	"RandomWallpaper/config"
	model "RandomWallpaper/models"
	"fmt"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSS struct {
}

var client *oss.Client

func init() {
	c, err := oss.New(config.Cfg.OSS.Endpoint, config.Cfg.OSS.AccessKeyID, config.Cfg.OSS.AccessKeySecret, oss.AuthVersion(oss.AuthV4), oss.Region("cn-hangzhou"))
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	client = c
}

func (OSS) UpLoad(image *model.Image, f multipart.File) error {
	bucket, err := client.Bucket(config.Cfg.OSS.Bucket)
	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}
	imgPath := image.Type + "/" + image.UpdateTime + "-" + image.Info
	err = bucket.PutObject(imgPath, f)
	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}
	image.Url = "https://callay-tc.oss-cn-hangzhou.aliyuncs.com/" + imgPath
	return nil
}
