package user_service

import (
	"fmt"
	"github.com/maxlcoder/gin-web/models/user_model"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (user *User) CreateUser() User {
	userData := map[string]interface{}{
		"name": user.Name,
	}
	userModel := user_model.Create(userData)
	fmt.Println(userModel)
	user.Id = userModel.ID
	return *user
}
