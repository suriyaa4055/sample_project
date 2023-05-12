package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	//dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "root"
	// dbName := "bookdb"

	dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "root"
	// dbName := "bookdb"

	db, err := sql.Open(dbDriver, "root:Suriyaa@tcp(127.0.0.1:3306)/demo")
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(mysql:3306)/"+dbName)
	if err != nil {
		//log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
	return db
}
