package models

import (
	"crypto/md5"
	"fmt"
	"github.com/Unknwon/com"
	"time"
)

const StatusEnable = 1
const StatusDisable = 0

type UserAuth struct {
	ID            int       `gorm:"primary_key"json:"id"`
	UserId        uint      `json:"user_id"`
	Identify      string    `json:"identify"`
	Password      string    `json:"-"`
	RememberToken string    `json:"-"`
	PasswordSalt  string    `json:"-"`
	Type          int       `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (UserAuth) TableName() string {
	return "user_auths"
}

func CheckUserAuthExist(identify string, _type int) bool {
	var userAuth UserAuth
	DB.Select("id").Where(UserAuth{Identify: identify, Type: _type}).First(&userAuth)
	if userAuth.ID > 0 {
		return true
	}
	return false
}

func DoRegister(identify string, _type int, password string, nickname string) error {

	tx := DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	user := User{Nickname: nickname, Status: StatusEnable}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	slat := com.RandomCreateBytes(10)

	userAuth := UserAuth{
		UserId:       user.ID,
		Identify:     identify,
		Password:     MakePassword(password, string(slat)),
		PasswordSalt: string(slat),
		Type:         _type,
	}

	if err := tx.Create(&userAuth).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func MakePassword(password string, slat string) string {
	data := []byte(fmt.Sprintf("%s%s", password, slat))
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
