package project

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository)Create(title *string, description *string) error{
	project := models.Project{
		Title: title,
		Description: description,
	}

	return r.DB.Create(&project).Error
}

func (r *Repository) Get(id string) (*models.Project, error) {
	project := new(models.Project)

	err := r.DB.Where("id = ?", id).Find(project).Error

	return project, err
}

func (r *Repository) GetAll() ([]*models.Project, error) {
	var projects []*models.Project

	err := r.DB.Find(&projects).Error

	return projects, err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

