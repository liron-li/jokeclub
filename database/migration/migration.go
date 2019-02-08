package migration

import (
	"jokeclub/app/models"
)

func Migration() {
	models.DB.AutoMigrate(&models.User{})
}
