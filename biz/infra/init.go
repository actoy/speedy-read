package infra

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
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

type Model struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func IdGenerate() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return node.Generate().Int64()
}
