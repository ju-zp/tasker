package project

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/ju-zp/tasker/svc/task"
	"strconv"
)

type Project models.Project

func FindByID(ID string, DB *gorm.DB) (*models.Project, []*models.TaskTodos){
	project, err := Find(ID, DB)

	if err != nil {
		fmt.Println("good luck son")
	}

	//_, err = task.FindByProjectID(ID, DB)

	_, err = task.CreateRepository(DB).FindByProjectId(project.ID)
	var taskTodos []*models.TaskTodos

	//if len(tasks) > 0 {
	//	for _, task := range tasks {
	//		todos, _ := todo.FindByTaskID(string(task.ID), DB)
	//		taskTodo := models.TaskTodos{
	//			Task:  task,
	//			Todos: todos,
	//		}
	//		taskTodos = append(taskTodos, &taskTodo)
	//	}
	//}

	return project, taskTodos
}

func Find(ID string, DB *gorm.DB) (*models.Project, error) {
	id, _ := strconv.Atoi(ID)

	project := models.Project{
		ID: int64(id),
	}

	err := DB.Find(&project).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}
