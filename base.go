package main

import (
	"github.com/haqim007/dairy_v0.1/config"
	"github.com/haqim007/dairy_v0.1/controllers"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/signup", inDB.CreateUser)
	router.POST("/login", inDB.GetUser)

	router.POST("/add_diary", inDB.CreateDiary)
	router.POST("/get_diary", inDB.GetDiaries)

	router.Run(":3000")
}
