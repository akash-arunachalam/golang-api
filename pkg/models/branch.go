package models

import (
	"fmt"
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

func GetBranchById(Id int64) (*Branch, *gorm.DB) {
	var getBranch Branch
	db := db.Where("ID = ?", Id).Find(&getBranch)
	fmt.Println(&getBranch)
	return &getBranch, db
}

func DeleteBranch(ID int64) Branch {
	var branch Branch
	db.Where("ID = ?", ID).Delete(branch)
	return branch
}
