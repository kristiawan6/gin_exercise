package helper

import (
	"gingonic/src/config"
	"gingonic/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Month{})
	config.DB.AutoMigrate(&models.Day{})
	config.DB.AutoMigrate(&models.Year{})

}
