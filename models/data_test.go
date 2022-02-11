package models

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	plainText := "test1"
	fmt.Println(Encrypt(plainText))
}
