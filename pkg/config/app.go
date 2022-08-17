package config

import (
	_ "database/sql"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	//d, err := gorm.Open("mysql", "akash:Akash#99@tcp(34.93.109.172:3306)/ak?charset=utf8&parseTime=True&loc=Local")
	dbDriver := "mysql"
	dbUser := "akash"
	dbPass := "Akash#99"
	dbInstance := "glassy-droplet-358909:asia-south1:billing"
	dbName := "ak"

	d, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@cloudsql("+dbInstance+")/"+dbName+"?parseTime=true")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
