package user

import (
	"context"
	"encoding/json"
	"fmt"
	"go-test/internal/domain/user/usecase"
	"net/http"
)

type dto struct {
	Id string `json:"id"`
}

type UserController struct {
	usecase *usecase.GetUserUsecase
}

func NewUserController(usecase *usecase.GetUserUsecase) UserController {
	return UserController{usecase: usecase}
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	var dto dto

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // Обработка ошибки
		return
	}

	fmt.Println(dto)

	u, err := c.usecase.GetById(context.Background(), dto.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func (c *UserController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/get-user-by-id", c.GetById)
}
