package task

import "context"

type IService interface {
	GetById(ctx context.Context, id string) (*TaskDomain, error)
}

type IRepository interface {
	GetById(ctx context.Context, id string) (*TaskDomain, error)
}
