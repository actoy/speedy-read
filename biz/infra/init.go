package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Init() {
	initMysql()
}

func initMysql() {
	dsn := "root:@Free4me@tcp(127.0.0.1:3306)/speedy_read?charset=utf8mb4&parseTime=True&loc=Local"
	openDBErr := error(nil)
	DB, openDBErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if openDBErr != nil {
		panic("mysql connect error")
	}
}
