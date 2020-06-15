package task

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(title *string, projectId int64) (*models.Task, error) {
	task := models.Task{
		Done:      false,
		ProjectID: projectId,
		Title:     title,
	}

	err := r.DB.Create(&task).Error

	return &task, err
}

func (r *Repository) Find(id string) (*models.Task, error) {
	task := models.Task{}

	err := r.DB.Where("id = ?", id).Find(&task).Error

	return &task, err
}

func (r *Repository) FindByProjectId(projectId int64) ([]*models.Task, error) {
	var tasks []*models.Task

	err := r.DB.Where("project_id = ?", projectId).Find(&tasks).Error

	return tasks, err
}

func (r *Repository) Delete(id string) error {
	task, err := r.Find(id)

	if err != nil {
		return err
	}

	err = r.DB.Delete(&task).Error

	return err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
