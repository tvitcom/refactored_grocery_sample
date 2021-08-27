package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tvitcom/refactored_grocery_sample/pkg/util"
)

type Shoppinglist struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Qty   int    `json:"qty" form:"qty" binding:"required,gte=0"`
	Unit  string `json:"unit" form:"unit" binding:"required"`
	Error error
}

func IndexWithPage(limit int, offset int) []Shoppinglist {
	db := util.GetDbConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shoppinglist LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()
	util.PanicError(err)

	var shoppingLists []Shoppinglist
	for rows.Next() {
		var sl Shoppinglist

		err = rows.Scan(&sl.Id, &sl.Name, &sl.Qty, &sl.Unit)
		shoppingLists = append(shoppingLists, sl)
	}

	return shoppingLists
}

func Count() int {
	db := util.GetDbConn()
	defer db.Close()

	var count int
	query := "SELECT COUNT(*) FROM shoppinglist"
	row := db.QueryRow(query)
	row.Scan(&count)
	return count
}

func Create(shoppingList Shoppinglist) (int64, error) {
	db := util.GetDbConn()
	defer db.Close()

	query := "INSERT INTO shoppinglist (name, qty, unit) VALUES(?, ?, ?);"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	res, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit)
	util.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	util.PanicError(getLastInsertIdErr)

	return id, queryErr
}

func FindById(id int64) Shoppinglist {
	var shoppingList Shoppinglist
	db := util.GetDbConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shoppinglist WHERE id = ?;"

	row := db.QueryRow(query, id)
	row.Scan(&shoppingList.Id, &shoppingList.Name, &shoppingList.Qty, &shoppingList.Unit)

	return shoppingList
}

func Put(id int64, shoppingList Shoppinglist) (Shoppinglist, error) {
	db := util.GetDbConn()
	defer db.Close()

	query := "UPDATE shoppinglist SET name = ?, qty = ?, unit = ? WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit, id)
	util.PanicError(queryErr)

	shoppingList.Id = id
	return shoppingList, queryErr
}

func Delete(shoppingList Shoppinglist) error {
	db := util.GetDbConn()
	defer db.Close()

	query := "DELETE FROM shoppinglist WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Id)
	util.PanicError(queryErr)

	return queryErr
}
