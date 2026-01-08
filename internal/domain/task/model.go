package task

import (
	"go-test/internal/domain/user"
	"time"
)

type TaskDomain struct {
	Id          string
	Name        string
	Status      string
	Responsible user.UserDomain
	CreatedAt   time.Time
}

func (t *TaskDomain) IsResponsible(id string) bool {
	return t.Responsible.Id == id
}
