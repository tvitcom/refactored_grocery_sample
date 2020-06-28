package shoppinglistrepo

import (
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/types"
	"github.com/dwahyudi/golang_grocery_sample/internal/app/grocery/util"
)

func Create(shoppingList types.ShoppingList) (int64, error) {
	db := util.DBConn()
	defer db.Close()

	query := "INSERT INTO shopping_list (name, qty, unit) VALUES(?, ?, ?);"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	res, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit)
	util.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	util.PanicError(getLastInsertIdErr)

	return id, queryErr
}

func FindById(id int64) (types.ShoppingList, error) {
	var shoppingList types.ShoppingList
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shopping_list WHERE id = ?;"

	row := db.QueryRow(query, id)
	row.Scan(&shoppingList.Id, &shoppingList.Name, &shoppingList.Qty, &shoppingList.Unit)

	return shoppingList, nil
}

func Put(id int64, shoppingList types.ShoppingList) (types.ShoppingList, error) {
	db := util.DBConn()
	defer db.Close()

	query := "UPDATE shopping_list SET name = ?, qty = ?, unit = ? WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit, id)
	util.PanicError(queryErr)

	shoppingList.Id = id
	return shoppingList, queryErr
}

func Delete(shoppingList types.ShoppingList) error {
	db := util.DBConn()
	defer db.Close()

	query := "DELETE FROM shopping_list WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Id)
	util.PanicError(queryErr)

	return queryErr
}
