package repository

import (
	"architecture.com/domain"
	"gorm.io/gorm"
)

type taskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(database *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		database: database,
	}
}

// Create implements domain.TaskRepository.
func (tr *taskRepository) Create(task *domain.Task) error {
	collection := tr.database.Create(&task).Error
	return collection
}

// FetchByUserID implements domain.TaskRepository.
func (tr *taskRepository) FetchByUserID(userID string) ([]domain.Task, error) {
	var tasks []domain.Task
	collection := tr.database.Where("user_id = ?", userID).Find(&tasks)
	return tasks, collection.Error
}
