package models

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(User{
		Name:     username,
		Password: password,
	}).First(&user)
	return user.ID > 0
}
