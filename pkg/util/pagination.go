package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"bgo/pkg/config"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * config.ConfigSetting.PageSize
	}

	return result
}
