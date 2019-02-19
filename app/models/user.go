package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key"json:"id"`
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

func (User) TableName() string {
	return "users"
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

func GetUserProfile(id uint) User {
	var user User
	DB.Where(User{ID: id}).First(&user)
	return user
}
