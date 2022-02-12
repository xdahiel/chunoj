package models

import (
	"crypto/sha1"
	"encoding/hex"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	IsActive bool   `json:"isActive"`
}

func Encrypt(plainText string) (cryptext string) {
	c := sha1.New()
	c.Write([]byte(plainText))
	bytes := c.Sum([]byte("chunojsalt"))
	return hex.EncodeToString(bytes)
}
