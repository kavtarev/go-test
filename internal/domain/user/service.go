package user

import (
	"context"
)

type UserService struct {
	repo IRepository
}

func NewUserService(repo IRepository) IService {
	return &UserService{repo: repo}
}

func (u *UserService) GetById(ctx context.Context, id string) (*UserDomain, error) {
	return u.repo.GetById(ctx, id)
}
