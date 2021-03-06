package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	MYSQL_SSH = "root:123/test?charset=utf8&parseTime=True&loc=Local"
)

func main() {
	db, err := gorm.Open(
		mysql.Open(MYSQL_SSH),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

}
