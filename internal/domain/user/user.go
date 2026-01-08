package user

import "context"

type IService interface {
	GetById(ctx context.Context, id string) (*UserDomain, error)
}

type IRepository interface {
	GetById(ctx context.Context, id string) (*UserDomain, error)
}
