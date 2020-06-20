package project

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/ju-zp/tasker/svc/task"
	"github.com/ju-zp/tasker/svc/todo"
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

func (r *Repository) Find(id string) (*models.Project, error) {
	project := new(models.Project)

	err := r.DB.Where("id = ?", id).Find(project).Error

	return project, err
}

func (r *Repository) FindAll() ([]*models.Project, error) {
	var projects []*models.Project

	err := r.DB.Find(&projects).Error

	return projects, err
}

func (r *Repository) FindByID(id string) (project *models.Project, taskTodos []*models.TaskTodos, err error) {
	project, err = r.Find(id)
	if err != nil {
		fmt.Println("error finding project")
	}

	taskRepository := task.CreateRepository(r.DB)

	tasks, err := taskRepository.FindByProjectId(project.ID)
	if err != nil {
		fmt.Println("error finding tasks")
	}

	if len(tasks) > 0 {
		todoRepository := todo.CreateRepository(r.DB)
		for _, task := range tasks {
			todos, _ := todoRepository.FindByTaskId(task.ID)
			taskTodo := models.TaskTodos{
				Task:  task,
				Todos: todos,
			}
			taskTodos = append(taskTodos, &taskTodo)
		}
	}

	return project, taskTodos, nil
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

