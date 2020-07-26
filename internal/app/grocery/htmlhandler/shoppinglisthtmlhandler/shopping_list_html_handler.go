package shoppinglisthtmlhandler

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/repo/shoppinglistrepo"
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := shoppinglistrepo.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "shopping_list/show.tmpl", shoppingList)
	}
}
