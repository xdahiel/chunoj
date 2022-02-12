package models

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	plainText := "test123"
	fmt.Println(Encrypt(plainText))
}
