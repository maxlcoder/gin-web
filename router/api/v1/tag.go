package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/models"
	"github.com/maxlcoder/gin-web/pkg/e"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"github.com/maxlcoder/gin-web/pkg/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	// 查询条件
	if name != "" {
		maps["tag_name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	log.Println(data)
	data["total"] = models.GetTotal(maps)

	code := e.SUCCESS
	message := e.GetMsg(code)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": data,
	})
}

func CreateTag(c *gin.Context)  {
	name := c.PostForm("name")
	fmt.Println(name)
}

func UpdateTag(c *gin.Context)  {

}

func DeleteTag(c *gin.Context)  {

}