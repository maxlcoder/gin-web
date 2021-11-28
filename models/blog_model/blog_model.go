package blog_model

import (
	"github.com/maxlcoder/gin-web/models"
	"gorm.io/gorm"
)

type BlogModel struct {
	models.Model
	Id        int
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"int"`
}

func GetBlogs(pageNum int, pageSize int, wheres interface{}) []*BlogModel {
	var blogs []*BlogModel
	err := models.DB.Model(&BlogModel{}).Find(&blogs)
	if err != nil {
		return nil
	}
	return blogs
}
