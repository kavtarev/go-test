package task

import "context"

type TaskService struct {
	repo IRepository
}

func (u *TaskService) GetById(ctx context.Context, id string) (*TaskDomain, error) {
	return u.repo.GetById(ctx, id)
}
