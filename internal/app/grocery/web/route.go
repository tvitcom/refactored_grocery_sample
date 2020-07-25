package web

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/htmlhandler/shoppinglisthtmlhandler"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("shopping-list/:id", shoppinglisthtmlhandler.ShowHandler)
	return r
}
