package domain

import (
	"gorm.io/gorm"
)

type Airline struct {
	gorm.Model
	Name string `json:"name"  gorm:"not null"`
	Logo string `json:"logo"`
}
