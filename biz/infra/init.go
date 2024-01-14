package infra

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	DB   *gorm.DB
	node *snowflake.Node
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
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func IdGenerate() int64 {
	return node.Generate().Int64()
}
