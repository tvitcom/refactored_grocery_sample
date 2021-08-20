package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/refactored_grocery_sample/internal/shoppinglist"
)

func Setup(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("shopping-list/", shoppinglist.IndexHandler)
	r.GET("shopping-list/show/:id", shoppinglist.ShowHandler)
	r.GET("shopping-list/new/", shoppinglist.NewHandler)
	r.POST("shopping-list/", shoppinglist.CreateHandler)
	r.GET("shopping-list/edit/:id", shoppinglist.EditHandler)
	r.POST("shopping-list/update/:id", shoppinglist.UpdateHandler)
	r.GET("shopping-list/delete/:id", shoppinglist.DeleteHandler)

	return r
}
