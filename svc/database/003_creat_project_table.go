package database

import "github.com/ju-zp/tasker/svc/services/tasker/models"

func (database Database)CreateProjectTable() {
	database.DB.AutoMigrate(&models.Project{})
}