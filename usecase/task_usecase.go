package usecase

import (
	"context"
	"time"

	"architecture.com/domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(tr domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: tr,
		contextTimeout: timeout,
	}
}

// Create implements domain.TaskUsecase.
func (tu *taskUsecase) Create(c context.Context, task *domain.Task) error {
	_, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(task)
}

// FetchByUserID implements domain.TaskUsecase.
func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	_, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(userID)
}
