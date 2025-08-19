package services

import (
	"context"
	"taskify/internal/store"
)

type TaskService struct {
	store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string, priority int32) (store.Task, error) {
	// Add validations and business logic
	task, err := s.store.CreateTask(ctx, title, description, priority)

	if err != nil {
		return store.Task{}, err
	}

	return task, err
}

func (s *TaskService) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	// Add validations and business logic

	task, err := s.store.GetTaskById(ctx, id)

	if err != nil {
		return store.Task{}, err
	}

	return task, err
}

func (s *TaskService) ListTasks(ctx context.Context) ([]store.Task, error) {
	// Add validations and business logic
	// Add page and limit

	task, err := s.store.ListTasks(ctx, 1, 50)

	if err != nil {
		return make([]store.Task, 0), err
	}

	return task, err
}

func (s *TaskService) UpdateTask(ctx context.Context, id int32, title, description string, priority int32) (store.Task, error) {
	// Add validations and business logic
	//
	task, err := s.store.UpdateTask(ctx, id, title, description, priority)

	if err != nil {
		return store.Task{}, err
	}

	return task, err
}

func (s *TaskService) DeleteTask(ctx context.Context, id int32) error {
	// Add validations and business logic
	//
	err := s.store.DeleteTask(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
