package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"suyuan/g"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * g.AppSetting.PageSize
	}

	return result
}
