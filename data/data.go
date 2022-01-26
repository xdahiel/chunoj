package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "dbname=chunoj user=root password=ycsv587hhh")
	if err != nil {
		log.Fatal(err)
	}
	return
}
