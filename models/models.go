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
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	dbName = setting.MysqlSetting.DB
	user = setting.MysqlSetting.User
	password = setting.MysqlSetting.Password
	host = setting.MysqlSetting.Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}
