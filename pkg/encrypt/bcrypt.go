package encrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	encode := string(hash)

	return encode
}

func BcryptCheck(encode string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(encode), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
