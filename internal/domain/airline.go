package domain

import (
	"gorm.io/gorm"
)

type Airline struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Logo string `json:"logo"`
}
