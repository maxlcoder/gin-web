package models

type Tag struct {
	Model
	TagName string `gorm:"default:'';not null" json:"name"`
	State   int    `gorm:"default:0;not null" json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTageByName(name string) bool {
	var tag Tag
	db.Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func CreateTag(name string) bool {
	db.Create(&Tag{
		TagName: name,
	})
	return true
}
