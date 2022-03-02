package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PasswordDigest string
}

const (
	PasswordCost = 12 //密码加密难度
)

// 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//校验密码
func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)) != nil
}