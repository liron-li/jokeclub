package migration

import (
	"goweb/app/models"
)

func Migration() {
	models.DB.AutoMigrate(&models.User{})
}
