package webhandler

import (
	"fmt"
	"github.com/dwahyudi/golang_grocery_sample/internal/repo"
	"github.com/dwahyudi/golang_grocery_sample/internal/types"
	"github.com/dwahyudi/golang_grocery_sample/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler(c *gin.Context) {
	var shoppingList types.ShoppingList

	// App level validation
	bindErr := c.BindJSON(&shoppingList)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := repo.Create(shoppingList)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		util.PanicError(insertErr)
	} else {
		shoppingList.Id = id
		c.JSON(http.StatusCreated, shoppingList)
	}
}

func ShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := repo.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, shoppingList)
	}
}

func PutHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	var shoppingList types.ShoppingList

	// App level validation
	bindErr := c.BindJSON(&shoppingList)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	// Check if resource exist
	foundShoppingList, _ := repo.FindById(id)
	if foundShoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Updating data
	shoppingList, err := repo.Put(foundShoppingList.Id, shoppingList)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, shoppingList)
	}
}

func DeleteHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)

	// Check if resource exist
	foundShoppingList, _ := repo.FindById(id)
	if foundShoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Deleting data
	err := repo.Delete(foundShoppingList)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusNoContent, "Successful Deletion")
	}
}
