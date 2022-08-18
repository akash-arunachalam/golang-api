package models

import (
	"fmt"
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var signinDb *gorm.DB

type Signin struct {
	gorm.Model

	Username string `gorm:""json:"username"`
	Password string `json:"password"`
}

type Userresponse struct {
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
	signinDb.AutoMigrate(&Signin{})
}

func (b *Signin) CreateUser() *Signin {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []Signin {
	var Users []Signin
	db.Find(&Users)
	return Users
}

func DeleteUser(ID int64) Signin {
	var user Signin
	db.Where("ID = ?", ID).Delete(user)
	return user
}

func ValidateUsername(Username string) (*Signin, *gorm.DB) {
	fmt.Println(Username)
	var userdetail Signin
	db := db.Where("username = ?", Username).Find(&userdetail)

	return &userdetail, db
}

func ValidateLogin(Username string) (*Signin, *gorm.DB) {
	fmt.Println("test" + Username)
	var userdetail Signin
	db := db.Where("username = 	?", Username).Find(&userdetail)

	return &userdetail, db
}
