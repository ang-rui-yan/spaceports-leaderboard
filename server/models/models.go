package models

import "gorm.io/gorm"

type Leaderboard struct {
	gorm.Model
	Username string `json:"username" gorm:"type:text;not null;default:null"`
	Score    int    `json:"score" gorm:"type:int;not null;default:null"`
}
