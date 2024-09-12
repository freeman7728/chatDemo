package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Sno    string `json:"sno" gorm:"unique"`
}
