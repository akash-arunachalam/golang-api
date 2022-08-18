package models

import (
	"fmt"
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var signinDb *gorm.DB

type User struct {
	gorm.Model

	Username string `gorm:""json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"token"`
}

func init() {
	config.Connect()
	signinDb = config.GetDB()
	signinDb.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID = ?", ID).Delete(user)
	return user
}

func ValidateUsername(Username string) (*User, *gorm.DB) {
	fmt.Println(Username)
	var userdetail User
	db := db.Where("username = ?", Username).Find(&userdetail)

	return &userdetail, db
}

func ValidateLogin(Username string) (*User, *gorm.DB) {
	fmt.Println("test" + Username)
	var userdetail User
	db := db.Where("username = 	?", Username).Find(&userdetail)

	return &userdetail, db
}

func GetUserById(Username string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("username = ?", Username).Find(&getUser)
	fmt.Println(&getUser)
	return &getUser, db
}
