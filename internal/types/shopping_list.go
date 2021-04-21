package types

type ShoppingList struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Qty   int    `json:"qty" form:"qty" binding:"required,gte=0"`
	Unit  string `json:"unit" form:"unit" binding:"required"`
	Error error
}
