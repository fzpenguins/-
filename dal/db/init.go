package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc/config"
	"grpc/pkg/constants"
	"log"
	"strings"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := strings.Join([]string{config.SqlUserName, ":", config.SqlPassword,
		"@tcp(" + config.MysqlIP + ")/", config.DataBase,
		"?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间

	err = Migrate()
	if err != nil {
		panic(err)
	}
}

func Migrate() (err error) {
	if err = DB.AutoMigrate(&Image{}); err != nil {
		log.Println("auto migrate failed", err)
		return
	}
	return
}
