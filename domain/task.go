package domain

import "context"

type Task struct {
	ID     uint   `bson:"_id" json:"-"`
	Title  string `bson:"title" form:"title" binding:"required" json:"title"`
	UserID uint   `bson:"userID" json:"-"`
}

type TaskRepository interface {
	Create(task *Task) error
	FetchByUserID(userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
