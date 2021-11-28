package tag_model

import "github.com/maxlcoder/gin-web/models"

type TagModel struct {
	models.Model
	Id int
	Name string
}

func GetTags() []TagModel {

	return []TagModel{}
}


