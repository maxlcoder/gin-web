package user_service

import (
	"fmt"
	"github.com/maxlcoder/gin-web/models/user_model"
)

type User struct {
	id int
	Name string
}

func (user *User) CreateUser() user_model.User {
	userData := map[string]interface{}{
		"name": user.Name,
	}
	userModel := user_model.Create(userData)
	fmt.Println(userModel)
	return userModel
}
