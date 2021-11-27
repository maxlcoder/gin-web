package models

type User struct {
	Model
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"default:'hello';not null" json:"username"`
	Password string `gorm:"default:'';not null" json:"password"`
}

func init() {
	db.AutoMigrate(&User{})
}

func (user User) Say() string {
	return user.Name
}

func (user *User) Bye() string {
	user.Name = "fdsfdfdsfdsfggg"
	return "Bye" + user.Name
}


