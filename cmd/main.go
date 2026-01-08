package main

import (
	"go-test/internal/app"
	"go.uber.org/fx"
)

func main() {
	application := fx.New(
		app.Init(),
	)

	application.Run()
}
