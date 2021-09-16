package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
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
