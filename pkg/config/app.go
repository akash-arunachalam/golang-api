package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	dbDriver := "mysql"
	dbUser := "akash"
	dbPass := "Akash#99"
	dbInstance := "glassy-droplet-358909:asia-south1:billing"
	dbName := "ak"

	d, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@unix(/cloudsql/"+dbInstance+")/"+dbName)
	if err != nil {
		fmt.Println("Connection Fail")

		panic(err)
	}
	fmt.Println("Connection Success")
	d.Query("CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)")
}

func GetDB() *gorm.DB {
	return db
}
