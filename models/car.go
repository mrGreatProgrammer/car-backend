package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Make     string  `json:"make"`
	CarModel string  `json:"carModel"`
	Year     int     `json:"year"`
	Price    float64 `json:"price"`
}
