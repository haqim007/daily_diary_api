package structs

import (
	"github.com/jinzhu/gorm"
)

//Login struct
type Login struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
