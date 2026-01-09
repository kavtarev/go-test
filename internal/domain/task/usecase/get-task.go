package usecase

import (
	"context"
	"go-test/internal/domain/task"
)

type GetTaskUsecase struct {
	service task.IService
}

func NewGetTaskUsecase(service task.IService) GetTaskUsecase {
	return GetTaskUsecase{service: service}
}

func (t *GetTaskUsecase) GetById(ctx context.Context, id string) (*task.TaskDomain, error) {
	return t.service.GetById(ctx, id)
}
