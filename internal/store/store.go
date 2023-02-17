package store

import "context"

type Task struct {
	ID          string
	Checked     bool
	Description string
}

type Store interface {
	StoreTask(ctx context.Context, task Task) error
	LoadTask(ctx context.Context, id string) (Task, error)
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context) ([]Task, error)
}
