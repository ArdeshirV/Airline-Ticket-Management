package domain

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name     string
	CityID   uint
	City     City
	Terminal string
}
