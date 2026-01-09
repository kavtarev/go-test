package app

import (
	"context"
	"fmt"
	"go-test/internal/presentation/task"
	"go-test/internal/presentation/user"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func NewMux(
	taskController task.TaskController,
	userController user.UserController,
) *http.ServeMux {
	mux := http.NewServeMux()
	taskController.RegisterRoutes(mux)
	userController.RegisterRoutes(mux)
	return mux
}

func ProvideDatabase(
	lc fx.Lifecycle,
) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", "postgres://postgres:postgres@localhost:5433/go_test?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("unable to open connection to database, error: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Add lifecycle hooks
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// First ping to check connection
			if err := db.PingContext(ctx); err != nil {
				return fmt.Errorf("unable to connect to database, error: %w", err)
			}
			log.Printf("Connected to DataBase")

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Printf("Closing database connections")
			return db.Close()
		},
	})

	return db, nil
}

func Init() fx.Option {
	return fx.Options(
		BuildDomainModule(),
		BuildInfrastructureModule(),
		BuildPresentationModule(),
		fx.Provide(NewMux, ProvideDatabase), fx.Invoke(func(mux *http.ServeMux) {
			log.Println("Starting server on :3000")
			if err := http.ListenAndServe(":3000", mux); err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
		}))
}
