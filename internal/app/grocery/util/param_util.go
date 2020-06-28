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
