package user_model

import (
	"fmt"
	"github.com/maxlcoder/gin-web/models"
	"gorm.io/gorm"
)

type User struct {
	models.Model
	ID        int    `gorm:"primary_key" json:"id"`
	Name      string `gorm:"default:'hello';not null" json:"username"`
	Password  string `gorm:"default:'';not null" json:"password"`
	Age       int `json:"age"`
	DeletedAt gorm.DeletedAt `gorm:"type:int(11);default:null" json:"password"`
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

func Create(params map[string]interface{}) User {
	user := User{
		Name: params["name"].(string),
	}
	result := models.DB.Create(&user);
	if result.Error != nil {
		fmt.Errorf("User Create Fail: %w \n", result.Error)
	}
	return user
}
