package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitMySQL() {
	var err error
	Db, err = sql.Open("postgres", "user=chunoj password=chunoj dbname=chunoj sslmode=disable")
	if err != nil {
		fmt.Println("can not connect to postgresql")
		panic(err)
	}
}
