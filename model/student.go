package model

import "github.com/jinzhu/gorm"

type Storable interface {
	gorm.Model
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Sno    string `json:"sno" gorm:"unique"`
}
