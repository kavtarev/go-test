package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go-test/internal/app"
	"go.uber.org/fx"
)

func main2() {
	application := fx.New(
		app.Init(),
	)

	application.Run()
	fmt.Println("should not be there")
}
