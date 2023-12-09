package service

import (
	"micgofiber/model"
	"micgofiber/repository"
	"strings"
)

type TodoService struct {
	TodoRepo *repository.TodoRepo
}

func NewTodoService(todoRepo *repository.TodoRepo) *TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
	}
}

func (tS *TodoService) GetData() []model.TodoModel {
	return tS.TodoRepo.Data
}

func (tS *TodoService) Action(request *model.TodoRequest) model.TodoResponse {
	err := model.TodoResponse{
		Title:   "Failed Action",
		Message: "action required",
	}
	switch strings.ToLower(request.Action) {
	case "add":
		err = tS.TodoRepo.AddTodo(request.Data)
	case "update":
		err = tS.TodoRepo.UpdateTodo(request.Data, request.Index)
	case "delete":
		err = tS.TodoRepo.DeleteTodo(request.Index)
	}
	return err
}
