package task

import (
	"context"
	"go-test/internal/domain/task/usecase"
	"net/http"
)

type TaskController struct {
	usecase *usecase.GetTaskUsecase
}

func NewTaskController(usecase *usecase.GetTaskUsecase) TaskController {
	return TaskController{usecase: usecase}
}

func (c *TaskController) GetById(w http.ResponseWriter, r *http.Request) {
	_, err := c.usecase.GetById(context.Background(), "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("task"))
}

func (c *TaskController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/get-task-by-id", c.GetById)
}
