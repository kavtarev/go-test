package app

import (
	taskDomain "go-test/internal/domain/task"
	taskDomainUsecase "go-test/internal/domain/task/usecase"

	userDomain "go-test/internal/domain/user"
	userDomainUsecase "go-test/internal/domain/user/usecase"

	"go.uber.org/fx"
)

func ProvideUserService(repo userDomain.IRepository) userDomain.IService {
	return userDomain.NewUserService(repo)
}

func ProvideUserGetByIdUsecase(service userDomain.IService) *userDomainUsecase.GetUserUsecase {
	return userDomainUsecase.NewGetUserUsecase(service)
}

func ProvideTaskService(repo taskDomain.IRepository) taskDomain.IService {
	return taskDomain.NewTaskService(repo)
}

func ProvideTaskGetByIdUsecase(service taskDomain.IService) *taskDomainUsecase.GetTaskUsecase {
	return taskDomainUsecase.NewGetTaskUsecase(service)
}

func ProvideDomain() fx.Option {
	return fx.Module("domain", fx.Provide(ProvideUserService, ProvideUserGetByIdUsecase, ProvideTaskService, ProvideTaskGetByIdUsecase))
}
