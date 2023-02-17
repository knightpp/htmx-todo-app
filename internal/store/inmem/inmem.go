package inmem

import (
	"context"
	"errors"
	"sync"
	"todo-htmx/internal/store"
)

var (
	ErrEmptyID = errors.New("error: empty id")
)

var _ store.Store = (*InMem)(nil)

type InMem struct {
	mu sync.Mutex
	m  map[string]store.Task
}

func New() *InMem {
	return &InMem{
		m: make(map[string]store.Task),
	}
}

func (m *InMem) StoreTask(_ context.Context, task store.Task) error {
	if task.ID == "" {
		return ErrEmptyID
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[task.ID] = task

	return nil
}

func (m *InMem) LoadTask(_ context.Context, id string) (store.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.m[id], nil
}

func (m *InMem) DeleteTask(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, id)

	return nil
}

func (m *InMem) ListTasks(_ context.Context) ([]store.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	tasks := make([]store.Task, 0, len(m.m))
	for _, v := range m.m {
		tasks = append(tasks, v)
	}

	return tasks, nil
}
