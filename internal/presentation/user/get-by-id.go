package user

import (
	"context"
	"go-test/internal/domain/user/usecase"
	"net/http"
)

type UserController struct {
	usecase usecase.GetUserUsecase
}

func NewUserController(usecase usecase.GetUserUsecase) UserController {
	return UserController{usecase: usecase}
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	_, err := c.usecase.GetById(context.Background(), "")
	if err != nil {
		w.Write([]byte("error"))
	}

	w.Write([]byte("user"))
}

func (c *UserController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/get-user-by-id", c.GetById)
}
