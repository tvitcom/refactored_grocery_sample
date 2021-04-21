package web

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/htmlhandler"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("shopping-list/show/:id", htmlhandler.ShowHandler)
	r.GET("shopping-list/new/", htmlhandler.NewHandler)
	r.POST("shopping-list/", htmlhandler.CreateHandler)
	r.GET("shopping-list/edit/:id", htmlhandler.EditHandler)
	r.POST("shopping-list/update/:id", htmlhandler.UpdateHandler)
	r.GET("shopping-list/", htmlhandler.IndexHandler)
	r.GET("shopping-list/delete/:id", htmlhandler.DeleteHandler)
	return r
}
