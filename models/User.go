package models

import (
	"chuns.cn/oj/databases"
	"log"
)

func GetUserByUsername(username string) (user User, err error) {
	err = databases.Db.QueryRow("select id, username, password from users where username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	return
}

func (user User) Create(username, password string) error {
	stmt, err := databases.Db.Prepare("insert into users (username, password) values ($1, $2) returning id, username, password")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = stmt.QueryRow(username, Encrypt(password)).Scan(&user.Id, &user.Username, &user.Password)
	return err
}
