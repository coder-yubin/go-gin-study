package util

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/coder-yubin/go-gin-study/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
