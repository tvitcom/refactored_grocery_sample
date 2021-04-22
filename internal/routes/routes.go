package routes

import (
	//shoppinglistapi "github.com/tvitcom/refactored_grocery_sample/internal/shoppinglist/api"
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/refactored_grocery_sample/internal/shoppinglist"
)

func Route(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// v1ShoppingList := r.Group("/v1/shopping-list")
	// {
	// 	v1ShoppingList.POST("/", shoppinglistapi.CreateHandler)
	// 	v1ShoppingList.GET("/:id", shoppinglistapi.ShowHandler)
	// 	v1ShoppingList.PUT("/:id", shoppinglistapi.PutHandler)
	// 	v1ShoppingList.DELETE("/:id", shoppinglistapi.DeleteHandler)
	// }

	r.GET("shopping-list/", shoppinglist.IndexHandler)
	r.GET("shopping-list/show/:id", shoppinglist.ShowHandler)
	r.GET("shopping-list/new/", shoppinglist.NewHandler)
	r.POST("shopping-list/", shoppinglist.CreateHandler)
	r.GET("shopping-list/edit/:id", shoppinglist.EditHandler)
	r.POST("shopping-list/update/:id", shoppinglist.UpdateHandler)
	r.GET("shopping-list/delete/:id", shoppinglist.DeleteHandler)

	return r
}
