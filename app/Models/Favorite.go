package Models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID int `json:"-"`
	WordID int `json:"-"`
	User   User
	Word   Word
}

type FavoriteAdd struct {
	UserID int `json:"user_id"`
	WordID int `json:"word_id"`
}
