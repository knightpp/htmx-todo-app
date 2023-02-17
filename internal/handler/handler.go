package handler

import (
	"errors"
	"fmt"
	"todo-htmx/internal/store"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID          string `form:"id"`
	Done        string `form:"done"`
	Description string `form:"description"`
}

func (t Task) toStoreTask() store.Task {
	return store.Task{
		ID:          t.ID,
		Checked:     t.Done == "on",
		Description: t.Description,
	}
}

type Handler struct {
	store store.Store
}

func New(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) Root(c echo.Context) error {
	tasks, err := h.store.ListTasks(c.Request().Context())
	if err != nil {
		return fmt.Errorf("list tasks: %w", err)
	}

	return c.Render(200, "root", map[string]any{"Tasks": tasks})
}

func (h *Handler) AddItem(c echo.Context) error {
	return c.Render(200, "item", store.Task{ID: uuid.NewString()})
}

func (h *Handler) PostItem(c echo.Context) error {
	task, err := h.bindTask(c)
	if err != nil {
		return err
	}

	storeTask := task.toStoreTask()

	err = h.store.StoreTask(c.Request().Context(), storeTask)
	if err != nil {
		return fmt.Errorf("store task: %w", err)
	}

	return c.Render(200, "item", storeTask)
}

func (h *Handler) DeleteItem(c echo.Context) error {
	task, err := h.bindTask(c)
	if err != nil {
		return err
	}

	err = h.store.DeleteTask(c.Request().Context(), task.ID)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return c.NoContent(200)
}

func (h *Handler) bindTask(c echo.Context) (Task, error) {
	var t Task
	err := c.Bind(&t)
	if err != nil {
		return t, fmt.Errorf("bind: %w", err)
	}
	if t.ID == "" {
		return t, errors.New("empty id")
	}

	return t, nil
}
