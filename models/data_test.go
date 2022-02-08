package models

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	plainText := "chunoj"
	fmt.Println(Encrypt(plainText))
}
