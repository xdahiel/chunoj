package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitMySQL() {
	var err error
	dsn := "chunoj:chunoj123@tcp(127.0.0.1:3306)/chunoj"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer Db.Close()
}
