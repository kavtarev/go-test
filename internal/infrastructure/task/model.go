package task

import (
	"database/sql"
	"go-test/internal/domain/task"
	"go-test/internal/domain/user"
	"time"
)

type TaskEntity struct {
	Id            string         `db:"id"`
	Name          string         `db:"name"`
	Status        string         `db:"status"`
	ResponsibleId sql.NullString `db:"responsible_id"`
	CreatedAt     time.Time      `db:"created_at"`
}

func (t *TaskEntity) ToDomain() *task.TaskDomain {
	return &task.TaskDomain{
		Id:          t.Id,
		Name:        t.Name,
		Status:      t.Status,
		Responsible: user.UserDomain{},
		CreatedAt:   t.CreatedAt,
	}
}
