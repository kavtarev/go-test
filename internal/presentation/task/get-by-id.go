package task

import (
	"context"
	"go-test/internal/domain/task/usecase"
	"net/http"
)

type TaskController struct {
	usecase usecase.GetTaskUsecase
}

func NewTaskController(usecase usecase.GetTaskUsecase) TaskController {
	return TaskController{usecase: usecase}
}

func (c *TaskController) GetById(w http.ResponseWriter, r *http.Request) {
	_, err := c.usecase.GetById(context.Background(), "")
	if err != nil {
		w.Write([]byte("error"))
	}

	w.Write([]byte("task"))
}
