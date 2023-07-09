package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string      `json:"username"`
	Password        string      `json:"password"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	RoleID          uint        `json:"roleid"`
	Role            Role        `json:"role"`
	Passengers      []Passenger `json:"passengers" gorm:"foreignKey:UserID"`
	IsLoginRequired bool        `json:"isloginrequired" gorm:"default:false"`
}
