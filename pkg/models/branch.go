package models

import (
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var branchDb *gorm.DB

type Branch struct {
	gorm.Model

	Branchname string `gorm:""json:"branchName"`
}

func init() {
	config.Connect()
	branchDb = config.GetDB()
	branchDb.AutoMigrate(&Branch{})
}

func (b *Branch) CreateBranch() *Branch {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBranch() []Branch {
	var Branches []Branch
	db.Find(&Branches)
	return Branches
}
