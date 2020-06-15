package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User struct
type User struct {
	gorm.Model
	Fullname string    `json:"fullname" form:"fullname" binding:"required"`
	Birthday time.Time `json:"birthday,time" form:"birthday,time" binding:"required"`
	Email    string    `json:"email" form:"email" binding:"required"`
	Username string    `json:"username" form:"username" binding:"required"`
	Password string    `json:"password" form:"password" binding:"required"`
}

//PostUser struct
type PostUser struct {
	gorm.Model
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Birthday string `json:"birthday,time" form:"birthday,time" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
