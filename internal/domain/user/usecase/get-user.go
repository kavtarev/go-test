package usecase

import (
	"context"
	"go-test/internal/domain/user"
)

type GetUserUsecase struct {
	service user.UserService
}

func (u *GetUserUsecase) GetById(ctx context.Context, id string) (*user.UserDomain, error) {
	return u.service.GetById(ctx, id)
}
