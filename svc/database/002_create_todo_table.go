package database

import (
	"github.com/ju-zp/tasker/svc/services/tasker/models"
)

func (database Database)CreateTodoTable() {
	database.DB.AutoMigrate(&models.Todo{})
	database.DB.Model(&models.Todo{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT")
}