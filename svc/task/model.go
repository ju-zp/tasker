package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
"github.com/ju-zp/tasker/svc/models"
"strconv"
)

type Task models.Task

func Find(ID string, DB *gorm.DB) (*models.Task, error) {
	id, _ := strconv.Atoi(ID)

	task := models.Task{
		ID: int64(id),
	}

	err := DB.Find(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func FindByProjectID(projectID string, DB *gorm.DB) ([]*models.Task, error) {
	var tasks []*models.Task

	fmt.Println(projectID)

	err := DB.Where("project_id = ?", projectID).Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}


