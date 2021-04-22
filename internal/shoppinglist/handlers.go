package shoppinglist

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/refactored_grocery_sample/internal/shoppinglist/model"
	"github.com/tvitcom/refactored_grocery_sample/internal/util"
	"net/http"
	"strconv"
)

func ShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := model.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "shoppinglist/show.tmpl", shoppingList)
	}
}

func NewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "shoppinglist/new.tmpl", gin.H{})
}

func CreateHandler(c *gin.Context) {
	var shoppingList model.Shoppinglist

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shoppinglist/new.tmpl", shoppingList)
		return
	}

	// Inserting data
	id, insertErr := model.Create(shoppingList)
	if insertErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(insertErr)
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/show/"+strconv.FormatInt(id, 10))
	}
}

func EditHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := model.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "shoppinglist/edit.tmpl", shoppingList)
}

func UpdateHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	var shoppingList model.Shoppinglist

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shoppinglist/edit.tmpl", shoppingList)
		return
	}

	foundShoppingList, _ := model.FindById(id)
	// Check if resource exist
	if foundShoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	}

	// Updating data
	shoppingList, updateErr := model.Put(foundShoppingList.Id, shoppingList)
	if updateErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(updateErr)
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/show/"+strconv.FormatInt(id, 10))
	}
}

func IndexHandler(c *gin.Context) {
	limit, offset, page := util.GetLimitOffset(c)
	shoppingLists := model.IndexWithPage(limit, offset)
	count := model.Count()
	pagination := util.ProcessPagination("shopping-list", count, page, limit)

	m := make(map[string]interface{})
	m["shoppingLists"] = shoppingLists
	m["pagination"] = pagination

	c.HTML(http.StatusOK, "shoppinglist/index.tmpl", m)
}

func DeleteHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	shoppingList, _ := model.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
		return
	}

	err := model.Delete(shoppingList)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		return
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/")
	}
}
