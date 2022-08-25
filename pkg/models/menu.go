package models

import (
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model

	Menutype   string `gorm:""json:"menuType"`
	Price      string `json:"price"`
	Itemlist   string `json:"itemList"`
	Sidedishes string `json:"sideDishes"`
	Menustatus bool   `json:"menuStatus"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Menu{})
}

func (b *Menu) CreateMenu() *Menu {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetMenu() []Menu {
	var Menu []Menu
	db.Find(&Menu)
	return Menu
}
