package util

import (
	"database/sql"
	"os"
)

func GetDbConn() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_GOLANG_LEARN_02"))
	PanicError(err)

	return db
}
