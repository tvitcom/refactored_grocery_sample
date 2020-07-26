package web

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/htmlhandler/shoppinglisthtmlhandler"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("shopping-list/show/:id", shoppinglisthtmlhandler.ShowHandler)
	r.GET("shopping-list/new/", shoppinglisthtmlhandler.NewHandler)
	r.POST("shopping-list/", shoppinglisthtmlhandler.CreateHandler)
	r.GET("shopping-list/edit/:id", shoppinglisthtmlhandler.EditHandler)
	r.POST("shopping-list/update/:id", shoppinglisthtmlhandler.UpdateHandler)
	r.GET("shopping-list/", shoppinglisthtmlhandler.IndexHandler)
	r.GET("shopping-list/delete/:id", shoppinglisthtmlhandler.DeleteHandler)
	return r
}
