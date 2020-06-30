package webapi

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/webhandler/shoppinglisthandler"
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
		v1ShoppingList.POST("/", shoppinglisthandler.CreateHandler)
		v1ShoppingList.GET("/:id", shoppinglisthandler.ShowHandler)
		v1ShoppingList.PUT("/:id", shoppinglisthandler.PutHandler)
		v1ShoppingList.DELETE("/:id", shoppinglisthandler.DeleteHandler)
	}

	return r
}
