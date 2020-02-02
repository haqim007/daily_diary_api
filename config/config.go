package config

import (
	"belajar_go_restapi/structs"
	"github.com/jinzhu/gorm"
)

//DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/gorestapidb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
