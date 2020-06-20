package todo

import (
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(content *string, taskID int64) error {
	todoStatus := false
	todo := models.Todo{
		Done: &todoStatus,
		Todo: content,
		TaskID: taskID,
	}

	return r.DB.Create(&todo).Error
}

func (r *Repository) Find(id string) (*models.Todo, error) {
	todo := models.Todo{}

	err := r.DB.Where("id = ?", id).Find(&todo).Error

	return &todo, err
}

func (r *Repository) FindByTaskId(taskId int64) ([]*models.Todo, error) {
	var todos []*models.Todo

	err := r.DB.Where("task_id = ?", taskId).Find(&todos).Error

	return todos, err
}

func (r *Repository) UpdateStatus(id string, status bool) error {
	todo, err := r.Find(id)

	if err != nil {
		return err
	}

	todo.Done = &status

	err = r.DB.Save(&todo).Error

	return err
}

func (r *Repository) DeleteByTodo(todo *models.Todo) error {
	err := r.DB.Delete(todo).Error

	return err
}

func (r *Repository) Delete(id string) error {
	todo, err := r.Find(id)

	if err != nil {
		return err
	}

	err = r.DeleteByTodo(todo)

	return err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}