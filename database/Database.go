package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func DatabaseInit() {
	var err error
	const MY_SQL = "root:@tcp(127.0.0.1:3306)/crud_karyawan?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MY_SQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("cannot connect to database!")
	}
	fmt.Println("Connected to database!")
}