package models

type User struct {
	Model
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func init() {
	db.Migrator().CreateTable(&User{})
}
