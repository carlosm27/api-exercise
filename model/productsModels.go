package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Sku            string  `json:"sku"`
	Name           string  `json:"name"`
	Brand          string  `json:"brand"`
	Price          float64 `json:"price"`
	Size           string  `json:"Size"`
	PrincipalImage string  `json:"image"`
}
