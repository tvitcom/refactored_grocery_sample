package types

type ShoppingList struct {
	Id   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
	Qty  int    `json:"qty" binding:"required,gte=0"`
	Unit string `json:"unit" binding:"required"`
}
