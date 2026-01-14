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
	var task TaskEntityWithResponsible

	err := r.db.GetContext(ctx, &task, `
		select
		    u.id as user_id
			, u.name as user_name
			, u.email as user_email
			, u.passport as user_passport
			, u.created_at as user_created_at
			, t.id
			, t.name
			, t.status
			, t.created_at
			, t.responsible_id
		from task t
		inner join users u on t.responsible_id=u.id`)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return task.ToDomain(), nil
}
