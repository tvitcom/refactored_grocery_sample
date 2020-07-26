package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInt64IdFromReqContext(c *gin.Context) int64 {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	return id
}

func GetLimitOffset(c *gin.Context) (int, int) {
	pageParam, _ := c.GetQuery("page")
	page, _ := strconv.Atoi(pageParam)
	if page == 0 {
		page = 1
	}

	limit := 10
	offset := limit * (page - 1)

	return limit, offset
}
