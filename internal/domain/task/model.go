package task

import (
	"go-test/internal/domain/user"
	"time"
)

type TaskDomain struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Status      string          `json:"status"`
	Responsible user.UserDomain `json:"responsible"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (t *TaskDomain) IsResponsible(id string) bool {
	return t.Responsible.Id == id
}
