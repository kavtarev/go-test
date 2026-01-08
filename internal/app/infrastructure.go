package app

import (
	taskDomain "go-test/internal/domain/task"
	taskInfra "go-test/internal/infrastructure/task"

	userDomain "go-test/internal/domain/user"
	userInfra "go-test/internal/infrastructure/user"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func ProvideTaskRepo(db *sqlx.DB) taskDomain.IRepository {
	return taskInfra.NewPostgresRepository(db)
}

func ProvideUserRepo(db *sqlx.DB) userDomain.IRepository {
	return userInfra.NewPostgresRepository(db)
}

func BuildInfrastructureModule() fx.Option {
	return fx.Module("infrastructure", fx.Provide(ProvideUserRepo, ProvideTaskRepo))
}
