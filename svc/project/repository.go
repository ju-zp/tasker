package project

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository)Create(title string) error{
	test := "here"
	project := models.Project{
		Title: &title,
		Description: &test,
	}

	return r.DB.Create(&project).Error
}

func (p *Repository) Get(id string) (*models.Project, error) {
	project := new(models.Project)

	err := p.DB.Where("id = ?", id).Find(project).Error

	return project, err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

