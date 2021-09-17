package Controllers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	//"github.com/go-playground/validator/v10"
)

func PasswordHash(password string) string {
	bs := []byte(password)
	hash, err2 := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	return string(hash)
}
func MatcherHash(hashPassword, password string) error {
	bs := []byte(password)
	hp := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hp, bs)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

type User struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=4,max=20"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}
