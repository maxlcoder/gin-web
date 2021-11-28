package blog_service

import "github.com/maxlcoder/gin-web/models/blog_model"

type Blog struct {
	id int
	name string
	content string

	pageNum   int
	pageSize int
}

func (blog *Blog) GetBlogs() ([]*blog_model.BlogModel, error)  {
	blog_model.GetBlogs(blog.pageNum, blog.pageSize, map[string]interface{} {

	})
}

func (blog *Blog) getWheres() map[string]interface{} {
	maps := make(map[string]interface{})

}
