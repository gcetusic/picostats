package main

import (
	"github.com/jinzhu/gorm"
)

// func GenerateAnonymousUser() sessionauth.User {
// 	return &User{}
// }

type User struct {
	gorm.Model
	Email         string `sql:"size:255" unique_index`
	Password      string `sql:"size:255"`
	Verified      bool   `sql:"not null"`
	authenticated bool
}

type Website struct {
	gorm.Model
	Owner   *User
	OwnerID int `sql:"index"`
}