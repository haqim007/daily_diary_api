package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Diary struct
type Diary struct {
	gorm.Model
	UserID  uint      `json:"user_id" form:"user_id"`
	Date    time.Time `json:"date,time" form:"date,time"`
	Content string    `json:"content" form:"content"`
}

//PostDiary struct for binding
type PostDiary struct {
	gorm.Model
	AccessToken string `json:"access_token" form:"access_token" binding:"required"`
	Date        string `json:"date,time" form:"date,time" binding:"required"`
	Content     string `json:"content" form:"content" binding:"required"`
}
