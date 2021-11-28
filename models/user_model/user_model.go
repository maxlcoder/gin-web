package user_model

import (
	"github.com/maxlcoder/gin-web/models"
	"gorm.io/gorm"
)

type User struct {
	models.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"default:'hello';not null" json:"username"`
	Password string `gorm:"default:'';not null" json:"password"`
	DeletedAt gorm.DeletedAt
}

// 自定注册
func init() {
}

func (user User) Say() string {
	return user.Name
}

func (user *User) Bye() string {
	user.Name = "fdsfdfdsfdsfggg"
	return "Bye" + user.Name
}


