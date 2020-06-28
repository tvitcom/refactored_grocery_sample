package webapi

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/webhandler/shopping_list_handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1ShoppingList := r.Group("/v1/shopping-list")
	{
		v1ShoppingList.POST("/", shopping_list_handler.CreateHandler)
		v1ShoppingList.GET("/:id", shopping_list_handler.ShowHandler)
		v1ShoppingList.PUT("/:id", shopping_list_handler.PutHandler)
		v1ShoppingList.DELETE("/:id", shopping_list_handler.DeleteHandler)
	}

	return r
}
