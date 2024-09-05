package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	ID             uint `gorm:"unique"`
	UserName       string
	PasswordDigest string
	Email          string //`gorm:"unique"`
	Avatar         string `gorm:"size:1000"`
	Phone          string
	Status         string
}

const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

func (u *User) SetPassword(password string) error {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(generateFromPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func (u *User) CheckUid() bool {
	result := DB.First(&User{}, "id = ?", u.ID)
	return result.RowsAffected > 0
}
