package models

import (
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var signinDb *gorm.DB

type User struct {
	gorm.Model

	Username string `gorm:""json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     string `json:"role"`
	Branch   int    `json:"branch"`
}

type Token struct {
	Message     string `json:"message"`
	TokenString string `json:"token"`
	Role        string `json:"role"`
	Branch      int    `json:"branch"`
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

	var userdetail User
	db := db.Where("username = ?", Username).Find(&userdetail)

	return &userdetail, db
}

func ValidateLogin(Username string) (*User, *gorm.DB) {

	var userdetail User
	db := db.Where("username = 	?", Username).Find(&userdetail)

	return &userdetail, db
}

func GetUserByName(Username string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("username = ?", Username).Find(&getUser)

	return &getUser, db
}
func GetUserById(ID int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID = ?", ID).Find(&getUser)

	return &getUser, db
}
