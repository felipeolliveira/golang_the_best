package pgstore

import (
	"context"
	"taskify/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	queries *Queries
	pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{queries: New(pool), pool: pool}
}

func (pgs *PGTaskStore) CreateTask(ctx context.Context, title string, description string, priority int32) (store.Task, error) {
	result, err := pgs.queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})

	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}, nil
}

func (pgs *PGTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	result, err := pgs.queries.GetTaskById(ctx, id)

	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}, nil
}

func (pgs *PGTaskStore) ListTasks(ctx context.Context, page int32, limit int32) ([]store.Task, error) {
	results, err := pgs.queries.ListTasks(ctx, ListTasksParams{
		Limit:  limit,
		Offset: (page - 1) * limit,
	})

	if err != nil {
		return []store.Task{}, err
	}

	list := make([]store.Task, 0, len(results))
	for _, result := range results {
		list = append(list, store.Task{
			Id:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Priority:    result.Priority,
			CreatedAt:   result.CreatedAt.Time,
			UpdatedAt:   result.UpdatedAt.Time,
		})
	}

	return list, nil
}

func (pgs *PGTaskStore) UpdateTask(ctx context.Context, id int32, title string, description string, priority int32) (store.Task, error) {
	result, err := pgs.queries.UpdateTask(ctx, UpdateTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	})

	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}, nil
}

func (pgs *PGTaskStore) DeleteTask(ctx context.Context, id int32) error {
	err := pgs.queries.DeleteTask(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
