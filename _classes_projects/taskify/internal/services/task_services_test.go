package services

import (
	"context"
	"fmt"
	"sort"
	"taskify/internal/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct {
	incrementPointer int32
	tasks            map[int32]store.Task
}

func NewMockTaskStore() MockTaskStore {
	return MockTaskStore{
		incrementPointer: 0,
		tasks:            make(map[int32]store.Task),
	}
}

func (mocktaskstore *MockTaskStore) CreateTask(ctx context.Context, title string, description string, priority int32) (store.Task, error) {
	newId := mocktaskstore.incrementPointer + 1
	newDate := time.Now()
	mocktaskstore.tasks[newId] = store.Task{
		Id:          newId,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   newDate,
		UpdatedAt:   newDate,
	}
	mocktaskstore.incrementPointer = newId

	return mocktaskstore.tasks[newId], nil
}

func (mocktaskstore *MockTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	task, ok := mocktaskstore.tasks[id]

	if !ok {
		return store.Task{}, fmt.Errorf("task not found")
	}

	return task, nil
}

func (mocktaskstore *MockTaskStore) ListTasks(ctx context.Context, page int32, limit int32) ([]store.Task, error) {
	list := make([]store.Task, 0, len(mocktaskstore.tasks))

	for _, task := range mocktaskstore.tasks {
		list = append(list, task)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})

	return list, nil
}

func (mocktaskstore *MockTaskStore) UpdateTask(ctx context.Context, id int32, title string, description string, priority int32) (store.Task, error) {
	task, err := mocktaskstore.GetTaskById(ctx, id)
	if err != nil {
		return task, err
	}

	mocktaskstore.tasks[id] = store.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	return mocktaskstore.tasks[id], nil
}
func (mocktaskstore *MockTaskStore) DeleteTask(ctx context.Context, id int32) error {
	_, err := mocktaskstore.GetTaskById(ctx, id)
	if err != nil {
		return err
	}

	delete(mocktaskstore.tasks, id)
	return nil
}

func TestCreateTask(t *testing.T) {
	// Arrange
	ctx := t.Context()
	mockStore := NewMockTaskStore()
	taskService := NewTaskService(&mockStore)

	// Act
	task, err := taskService.CreateTask(ctx, "Mock Test Task", "Mock Test Description", 1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTask(t *testing.T) {
	ctx := t.Context()
	mockStore := NewMockTaskStore()
	taskService := NewTaskService(&mockStore)

	_, err := taskService.CreateTask(ctx, "task 1", "Mock Test Description", 1)
	_, err = taskService.CreateTask(ctx, "task 2", "Mock Test Description", 10)
	_, err = taskService.CreateTask(ctx, "task 3", "Mock Test Description", 20)
	assert.NoError(t, err)

	task, err := taskService.GetTaskById(ctx, 2)

	assert.NoError(t, err)
	assert.Equal(t, "task 2", task.Title)
	assert.Equal(t, "Mock Test Description", task.Description)
	assert.Equal(t, int32(10), task.Priority)
}

func TestListTasks(t *testing.T) {
	ctx := t.Context()
	mockStore := NewMockTaskStore()
	taskService := NewTaskService(&mockStore)

	_, err := taskService.CreateTask(ctx, "task 1", "Mock Test Description", 1)
	_, err = taskService.CreateTask(ctx, "task 2", "Mock Test Description", 10)
	_, err = taskService.CreateTask(ctx, "task 3", "Mock Test Description", 20)
	assert.NoError(t, err)

	tasks, err := taskService.ListTasks(ctx)

	assert.NoError(t, err)
	assert.Len(t, tasks, 3)

	fistTask := tasks[2]
	assert.Equal(t, "task 1", fistTask.Title)
}

func TestUpdateTasks(t *testing.T) {
	ctx := t.Context()
	mockStore := NewMockTaskStore()
	taskService := NewTaskService(&mockStore)

	_, err := taskService.CreateTask(ctx, "task 1", "Mock Test Description", 1)
	_, err = taskService.CreateTask(ctx, "task 2", "Mock Test Description", 10)
	_, err = taskService.CreateTask(ctx, "task 3", "Mock Test Description", 20)
	assert.NoError(t, err)

	firstTask, err := taskService.GetTaskById(ctx, 1)
	secondTask, err := taskService.GetTaskById(ctx, 2)
	time.Sleep(10 * time.Millisecond)
	updatedTask, err := taskService.UpdateTask(ctx, 3, "task 3 - Edit", "Updated Mock Test Description", 25)

	assert.NoError(t, err)
	assert.Equal(t, "task 3 - Edit", updatedTask.Title)
	assert.Equal(t, "Updated Mock Test Description", updatedTask.Description)

	assert.Equal(t, firstTask.CreatedAt, firstTask.UpdatedAt)
	assert.Equal(t, secondTask.CreatedAt, secondTask.UpdatedAt)
	assert.NotEqual(t, updatedTask.CreatedAt, updatedTask.UpdatedAt)
	assert.True(t, updatedTask.CreatedAt.Before(updatedTask.UpdatedAt))
}
