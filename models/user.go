package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	Model
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"index;size:100;default:'hello';not null;comment:姓名"`
	Password  string `gorm:"size:300;default:'';not null"`
	Age       int8 `gorm:"default:0;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:int(11);default:null"`
}

func GetUserById(id int) (User, error) {
	var user User
	result := db.Where("id", id).First(&user)
	return user, result.Error
}

func CreateUser(params map[string]interface{}) User {
	user := User{
		Name: params["name"].(string),
	}
	result := db.Create(&user);
	if result.Error != nil {
		fmt.Errorf("User Create Fail: %w \n", result.Error)
	}
	return user
}
