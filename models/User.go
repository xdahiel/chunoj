package models

import (
	"chuns.cn/oj/databases"
	"log"
)

func GetUserByUsername(username string) (user User, err error) {
	err = databases.Db.QueryRow("select id, username, password, isactive, email from users where username = $1", username).Scan(&user.Id, &user.Username, &user.Password, &user.IsActive, &user.Email)
	return
}

func GetUserByEmail(email string) (user User, err error) {
	err = databases.Db.QueryRow("select id, username, password, isactive, email from users where email=$1", email).Scan(&user.Id, &user.Username, &user.Password, &user.IsActive, &user.Email)
	return
}

func (user User) Create() error {
	stmt, err := databases.Db.Prepare("insert into users (username, password, email, isactive) values ($1, $2, $3, false) returning id, username, password, isactive, email")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = stmt.QueryRow(user.Username, Encrypt(user.Password), user.Email).Scan(&user.Id, &user.Username, &user.Password, &user.IsActive, &user.Email)
	return err
}
