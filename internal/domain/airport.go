package domain

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name     string `json:"name"`
	CityID   uint   `json:"cityid"`
	City     City   `json:"city"`
	Terminal string `json:"terminal"`
}
