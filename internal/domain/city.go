package domain

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name string `json:"name"`
}
