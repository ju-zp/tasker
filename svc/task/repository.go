package task

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(title *string, projectId int64) error{
	task := &models.Task{
		Done:      false,
		ProjectID: projectId,
		Title:     title,
	}

	err := r.DB.Create(task).Error

	return err
}


func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
