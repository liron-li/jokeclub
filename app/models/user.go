package models

import "time"

type User struct {
	ID              uint       `gorm:"primary_key"json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	EmailVerifiedAt string     `json:"-"`
	Password        string     `json:"-"`
	Avatar          string     `json:"avatar"`
	Slogan          string     `json:"slogan"`
	Status          string     `json:"-"`
	RememberToken   string     `json:"-"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"-"`
	DeletedAt       *time.Time `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func CheckAuth(username, password string) bool {
	var user User
	DB.Select("id").Where(User{Name: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
