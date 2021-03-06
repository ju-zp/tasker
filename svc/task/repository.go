package task

import (
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/services/tasker/models"
	"github.com/ju-zp/tasker/svc/todo"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(title *string, projectId int64) (*models.Task, error) {
	task := models.Task{
		Done:      swag.Bool(false),
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

func (r *Repository) DeleteByTask(task *models.Task) error {
	todoRepo := todo.CreateRepository(r.DB)

	todos, err := todoRepo.FindByTaskId(task.ID)

	if err != nil {
		return err
	}

	for _, todo := range todos {
		err := todoRepo.DeleteByTodo(todo)
		if err != nil {
			return nil
		}
	}

	err = r.DB.Delete(&task).Error

	return err
}

func (r *Repository) Delete(id string) error {
	task, err := r.Find(id)

	if err != nil {
		return err
	}

	err = r.DeleteByTask(task)

	return err
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
