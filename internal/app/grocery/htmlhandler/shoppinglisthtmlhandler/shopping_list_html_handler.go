package shoppinglisthtmlhandler

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/repo/shoppinglistrepo"
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/types"
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func NewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "shopping_list/new.tmpl", gin.H{})
}

func CreateHandler(c *gin.Context) {
	var shoppingList types.ShoppingList

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shopping_list/new.tmpl", shoppingList)
		return
	}

	// Inserting data
	id, insertErr := shoppinglistrepo.Create(shoppingList)
	if insertErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(insertErr)
	} else {
		c.Redirect(http.StatusFound, "show/"+strconv.FormatInt(id, 10))
	}
}

func EditHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := shoppinglistrepo.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "shopping_list/edit.tmpl", shoppingList)
}

func UpdateHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	var shoppingList types.ShoppingList

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shopping_list/edit.tmpl", shoppingList)
		return
	}

	foundShoppingList, _ := shoppinglistrepo.FindById(id)
	// Check if resource exist
	if foundShoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	}

	// Updating data
	shoppingList, updateErr := shoppinglistrepo.Put(foundShoppingList.Id, shoppingList)
	if updateErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(updateErr)
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/show/"+strconv.FormatInt(id, 10))
	}
}
