package models

import "time"

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

func CheckAuth(username, password string) bool {
	//var user User
	//DB.Select("id").Where(User{Nickname: username, Password: password}).First(&user)
	//if user.ID > 0 {
	//	return true
	//}

	return false
}

func CheckUserExist(user User) bool {
	var _user User
	DB.Select("id").Where(user).First(&_user)
	if _user.ID > 0 {
		return true
	}
	return false
}
