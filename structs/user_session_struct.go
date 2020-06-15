package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

//UserSession struct
type UserSession struct {
	gorm.Model
	UserID       uint
	Token        string
	RefreshToken string
	IssuedAt     time.Time
}
