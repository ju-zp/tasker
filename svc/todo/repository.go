package todo

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository)Create(content *string, taskID int64) error{
	todoStatus := false
	todo := models.Todo{
		Done: &todoStatus,
		Todo: content,
		TaskID: taskID,
	}

	return r.DB.Create(&todo).Error
}

func (r *Repository) Get(id string) (*models.Todo, error) {
	todo := models.Todo{}

	err := r.DB.Where("id = ?", id).Find(&todo).Error

	return &todo, err
}

func (r *Repository) GetAll() ([]*models.Todo, error) {
	var todos []*models.Todo

	err := r.DB.Find(&todos).Error

	return todos, err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}