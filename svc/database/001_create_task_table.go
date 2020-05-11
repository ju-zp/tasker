package database

import (
	"github.com/ju-zp/tasker/svc/models"
)

func (database Database)CreateTaskTable() {
	exists := database.DB.HasTable(&models.Task{})
	if !exists {
		database.DB.AutoMigrate(&models.Task{})
	}
}