package models

type User struct {
	ID int `gorm:"primary_key" json:"id"`
	Name string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(User{
		Name: username,
		Password: password,
	}).First(&user)
	return user.ID > 0
}