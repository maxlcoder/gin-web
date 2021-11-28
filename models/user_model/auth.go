package user_model

import "github.com/maxlcoder/gin-web/models"

func CheckAuth(username, password string) bool {
	var user User
	models.DB.Select("id").Where(User{
		Name:     username,
		Password: password,
	}).First(&user)
	return user.ID > 0
}
