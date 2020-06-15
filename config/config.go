package config

import (
	"fmt"
	"log"
	"os"

	"github.com/haqim007/dairy_v0.1/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//DBInit create connection to database
func DBInit() *gorm.DB {
	var errEnv error
	errEnv = godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error getting env, not comming through %v", errEnv)
	} else {
		fmt.Println("We are getting the env values")
	}
	// fmt.Println("host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" sslmode=disable port="+os.Getenv("DB_PORT")+" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME")+" password="+os.Getenv("DB_PASSWORD"))
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.User{}, structs.UserSession{}, structs.Diary{})

	return db
}
