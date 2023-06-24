package Models

import "gorm.io/gorm"

type Word struct {
	Word        string `gorm:"not null" validate:"required" json:"Word"`
	Translation string `gorm:"not null" validate:"required" json:"Translation"`
	Example     string `gorm:"not null" validate:"required" json:"Example"`
	gorm.Model
}
