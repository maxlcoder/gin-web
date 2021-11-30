package models

type Tag struct {
	Model
	Id int
	Name string
}

func GetTags() []Tag {

	return []Tag{}
}
