package app

import (
	userUsecase "go-test/internal/domain/user/usecase"
	userPresentation "go-test/internal/presentation/user"

	taskUsecase "go-test/internal/domain/task/usecase"
	taskPresentation "go-test/internal/presentation/task"

	"go.uber.org/fx"
)

func ProvideGetUserByIdHandler(u *userUsecase.GetUserUsecase) userPresentation.UserController {
	return userPresentation.NewUserController(u)
}

func ProvideGetTaskByIdHandler(u *taskUsecase.GetTaskUsecase) taskPresentation.TaskController {
	return taskPresentation.NewTaskController(u)
}

func BuildPresentationModule() fx.Option {
	return fx.Module(
		"presentation",
		fx.Provide(
			ProvideGetUserByIdHandler,
			ProvideGetTaskByIdHandler,
		),
	)
}
