package models

import (
	"crypto/md5"
	"fmt"
	"github.com/Unknwon/com"
	"time"
)

const (
	StatusEnable  = 1
	StatusDisable = 0
)

type User struct {
	ID        int        `gorm:"primary_key"json:"id"`
	Nickname  string     `json:"nickname"`
	Email     string     `gorm:"default:null"json:"email"`
	PhoneArea string     `gorm:"default:'86'"json:"phone_area"`
	Phone     string     `gorm:"default:null"json:"phone"`
	Avatar    string     `json:"avatar"`
	Slogan    string     `json:"slogan"`
	Status    int        `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type UserAuth struct {
	ID            int       `gorm:"primary_key"json:"id"`
	UserId        int       `json:"user_id"`
	Identify      string    `json:"identify"`
	Password      string    `json:"-"`
	RememberToken string    `json:"-"`
	PasswordSalt  string    `json:"-"`
	Type          int       `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (UserAuth) TableName() string {
	return "user_auths"
}

func CheckAuth(username, password string) (UserAuth, bool) {

	var userAuth UserAuth
	DB.Where(UserAuth{Identify: username}).First(&userAuth)

	if userAuth.ID > 0 && MakePasswordHash(password, userAuth.PasswordSalt) == userAuth.Password {
		return userAuth, true
	}

	return userAuth, false
}

func CheckUserExist(user User) bool {
	var userModel User
	DB.Select("id").Where(user).First(&userModel)
	if userModel.ID > 0 {
		return true
	}
	return false
}

func GetUserProfile(id int) User {
	var user User
	DB.Where(User{ID: id}).First(&user)
	return user
}

func CheckUserAuthExist(identify string, typeValue int) bool {
	var userAuth UserAuth
	DB.Select("id").Where(UserAuth{Identify: identify, Type: typeValue}).First(&userAuth)
	if userAuth.ID > 0 {
		return true
	}
	return false
}

func DoRegister(identify string, typeValue int, password string, nickname string) error {

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
		Password:     MakePasswordHash(password, string(slat)),
		PasswordSalt: string(slat),
		Type:         typeValue,
	}

	if err := tx.Create(&userAuth).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func MakePasswordHash(password string, slat string) string {
	data := []byte(fmt.Sprintf("%s%s", password, slat))
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
