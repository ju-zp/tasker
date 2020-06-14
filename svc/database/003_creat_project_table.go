package database

import "github.com/ju-zp/tasker/svc/models"

func (database Database)CreateProjectTable() {
	database.DB.AutoMigrate(&models.Project{})
}