package models

import (
	"simple-REST-master/pkg/config"

	"github.com/jinzhu/gorm"
)

var signinDb *gorm.DB

type Signin struct {
	gorm.Model

	Username string `gorm:""json:"username"`
	Password string `json:"password"`
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
