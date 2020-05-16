package database

import (
	"github.com/ju-zp/tasker/svc/models"
)

func (database Database)CreateTaskTable() {
	database.DB.AutoMigrate(&models.Task{})
}