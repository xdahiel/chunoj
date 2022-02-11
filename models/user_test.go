package models

import (
	"chuns.cn/oj/databases"
	"fmt"
	"log"
	"testing"
)

func init() {
	databases.InitMySQL()
}

func TestUser_Create(t *testing.T) {
	user := new(User)
	err := user.Create("test100", "test100")
	if err != nil {
		log.Fatalln(err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	user, err := GetUserByUsername("test100")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
}
