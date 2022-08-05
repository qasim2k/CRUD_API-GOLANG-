package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:@tcp(localhost:3306)/cafedb?parseTime=true"
var err error

func DataMiration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(("connection failed"))
	}
	Database.AutoMigrate(&Cafe{})
}
