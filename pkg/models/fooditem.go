package models

import (
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

type FoodItem struct {
	gorm.Model

	Itemname     string `gorm:""json:"itemName"`
	Itemprice    string `json:"itemPrice"`
	Sidedishes   string `json:"sideDishes"`
	Itemquantity string `json:"itemQuantity"`
	Itemstatus   bool   `json:"itemStatus"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&FoodItem{})
}

func (b *FoodItem) CreateFooditem() *FoodItem {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetFooditems() []FoodItem {
	var Fooditems []FoodItem
	db.Find(&Fooditems)
	return Fooditems
}
