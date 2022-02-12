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
	user := User{
		Username: "test123",
		Password: "test123",
		Email:    "123@cc.com",
	}
	err := user.Create()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	user, err := GetUserByUsername("test123")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
}
