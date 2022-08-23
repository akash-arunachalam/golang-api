package models

import (
	"fmt"
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

type SideDish struct {
	gorm.Model
	//Id          string `json:"id"`
	Item     string `gorm:""json:"item"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&SideDish{})
}

func (b *SideDish) CreateSidedish() *SideDish {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetSidedishes() []SideDish {
	var Sidedishes []SideDish
	db.Find(&Sidedishes)
	return Sidedishes
}

func GetSidedishById(Id int64) (*SideDish, *gorm.DB) {
	var getSidedish SideDish
	db := db.Where("ID = ?", Id).Find(&getSidedish)
	fmt.Println(&getSidedish)
	return &getSidedish, db
}

func DeleteSidedish(ID int64) SideDish {
	var sidedish SideDish
	db.Where("ID = ?", ID).Delete(sidedish)
	return sidedish
}
