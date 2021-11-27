package models

import (
	"fmt"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `gorm:"autoCreateTime;not null" json:"created_at"`
	UpdatedAt int `gorm:"autoUpdateTime;not null" json:"updated_at"`
}

func init() {
	var (
		err                          error
		dbName, user, password, host, port string
	)

	dbName = setting.MysqlSetting.DB
	user = setting.MysqlSetting.User
	password = setting.MysqlSetting.Password
	host = setting.MysqlSetting.Host
	port = setting.MysqlSetting.Port

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}
