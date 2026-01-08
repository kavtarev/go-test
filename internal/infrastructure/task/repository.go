package task

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go-test/internal/domain/task"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(
	db *sqlx.DB,
) task.IRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetById(ctx context.Context, id string) (*task.TaskDomain, error) {
	var task TaskEntity

	err := r.db.GetContext(ctx, &task, ``)
	if err != nil {
		return nil, fmt.Errorf("failed to count systems: %w", err)
	}

	return task.ToDomain(), nil
}
