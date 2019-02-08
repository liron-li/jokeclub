package models

type User struct {
	Model
	Username string `gorm:"type:varchar(32);unique_index;default:''"`
	Password string `gorm:"type:varchar(50);default:''"`
}

func CheckAuth(username, password string) bool {
	var user User
	DB.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
