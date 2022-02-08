package models

import "chuns.cn/oj/databases"

func GetUserByUsername(username string) (user User, err error) {
	err = databases.Db.QueryRow("select id, username, password from user where username = $1", username).Scan(user.Id, user.Username, user.Password)
	return
}

//func (user User) CreateUser(username, password string) error {
//	stmt, err := databases.Db.Prepare("insert into user(username, password) values($1, $2)")
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//}
