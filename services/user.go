package services

import (
	"github.com/maxlcoder/gin-web/models"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (user *User) GetUser() (*User, error) {
	userModel, err := models.GetUserById(user.Id)
	if err != nil {
		return nil, err
	}
	user.Name = userModel.Name
	return user, nil
}

func (user *User) CreateUser() User {
	userData := map[string]interface{}{
		"name": user.Name,
	}
	userModel := models.CreateUser(userData)
	user.Id = userModel.ID
	return *user
}
