package task

import "context"

type TaskService struct {
	repo IRepository
}

func NewTaskService(repo IRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (u *TaskService) GetById(ctx context.Context, id string) (*TaskDomain, error) {
	return u.repo.GetById(ctx, id)
}
