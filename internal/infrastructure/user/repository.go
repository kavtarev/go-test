package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go-test/internal/domain/user"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(
	db *sqlx.DB,
) user.IRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetById(ctx context.Context, id string) (*user.UserDomain, error) {
	var user UserEntity

	err := r.db.GetContext(ctx, &user, `select * from user`)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return user.ToDomain(), nil
}
