package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Id         int    `json:"id" gorm:"primaryKey"`
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer string `json:"answer" gorm:"text;not null;default:null`
}
