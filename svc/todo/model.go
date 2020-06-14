package todo

import (
"github.com/jinzhu/gorm"
"github.com/ju-zp/tasker/svc/models"
"strconv"
)

type Task models.Task

func Find(ID string, DB *gorm.DB) (*models.Todo, error) {
	id, _ := strconv.Atoi(ID)

	todo := models.Todo{
		ID: int64(id),
	}

	err := DB.Find(&todo).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func FindByTaskID(taskID string, DB *gorm.DB) ([]*models.Todo, error) {
	var todos []*models.Todo

	err := DB.Where("task_id = ?", taskID).Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil
}