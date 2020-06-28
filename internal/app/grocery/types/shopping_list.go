package types

type ShoppingList struct {
	Id   int64  `json:"id"`
	Name string `json:"name" binding:"required,omitempty"`
	Qty  int    `json:"qty" binding:"required,gte=0,omitempty"`
	Unit string `json:"unit" binding:"required,omitempty"`
}
