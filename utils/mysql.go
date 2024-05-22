package utils

import (
	"RandomWallpaper/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Cfg.Mysql.Username, config.Cfg.Mysql.Password, config.Cfg.Mysql.Host, config.Cfg.Mysql.Port, config.Cfg.Mysql.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
