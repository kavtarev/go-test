package user

import (
	"context"
	"go-test/internal/domain/user/usecase"
	"net/http"
)

type UserController struct {
	usecase usecase.GetUserUsecase
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	_, err := c.usecase.GetById(context.Background(), "")
	if err != nil {
		w.Write([]byte("error"))
	}

	w.Write([]byte("user"))
}
