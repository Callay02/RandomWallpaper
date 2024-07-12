package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	OSS    OSS    `yaml:"oss"`
}

type Server struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	ImgSavePath string `yaml:"imgsavepath"`
	Token       bool   `yaml:"token"`
}

type Mysql struct {
	DbName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type OSS struct {
	Endpoint        string `yaml:"endpoint"`
	Bucket          string `yaml:"bucket"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}

var Cfg Config

func init() {
	// 打开 YAML 文件
	file, err := os.Open("./config/config.yml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&Cfg)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}
}
