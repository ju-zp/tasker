package database

import (
	"github.com/ju-zp/tasker/svc/services/tasker/models"
)

func (database Database)CreateTaskTable() {
	database.DB.AutoMigrate(&models.Task{})
}