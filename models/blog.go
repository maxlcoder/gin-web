package models

import "gorm.io/gorm"

type Blog struct {
	Model
	Id        int
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"int"`
}

func GetBlogs(pageNum int, pageSize int, wheres interface{}) []*Blog {
	var blogs []*Blog
	err := db.Model(&Blog{}).Find(&blogs)
	if err != nil {
		return nil
	}
	return blogs
}

