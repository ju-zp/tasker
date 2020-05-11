package database

import (
	"github.com/ju-zp/tasker/svc/models"
)

func (database Database)CreateTodoTable() {
	exists := database.DB.HasTable(&models.Todo{})
	if !exists {
		database.DB.AutoMigrate(&models.Todo{})
		database.DB.Model(&models.Todo{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT")
	}
}