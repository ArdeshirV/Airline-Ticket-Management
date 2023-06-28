package domain

import (
	"time"
)

type User struct {
	ID              int         `json:"id"`
	Username        string      `json:"username"`
	Password        string      `json:"password"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	CreatedAt       time.Time   `json:"createdat"`
	RoleID          uint        `json:"roleid"`
	Role            Role        `json:"role"`
	Passengers      []Passenger `json:"passengers" gorm:"foreignKey:UserID"`
	IsLoginRequired bool        `json:"isloginrequired" gorm:"default:false"`
}
