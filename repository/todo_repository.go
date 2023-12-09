package repository

import (
	"micgofiber/model"
)

type TodoRepo struct {
	Data []model.TodoModel
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{}
}

func (tR *TodoRepo) AddTodo(newTodo model.TodoModel) model.TodoResponse {
	tR.Data = append(tR.Data, newTodo)
	return model.TodoResponse{
		Title: "Success add todo",
	}
}

func (tR *TodoRepo) FindTodo(index int, action string) model.TodoResponse {
	for i := range tR.Data {
		if i == index {
			return model.TodoResponse{}
		}
	}
	return model.TodoResponse{
		Title:   action + " Failed",
		Message: "Data Not Found",
	}
}

func (tR *TodoRepo) UpdateTodo(newTodo model.TodoModel, index int) model.TodoResponse {
	err := tR.FindTodo(index, "Update")
	if len(err.Message) == 0 {
		tR.Data[index] = newTodo
		err.Title = "Success update todo"
	}
	return err
}

func (tR *TodoRepo) DeleteTodo(index int) model.TodoResponse {
	err := tR.FindTodo(index, "Delete")
	if len(err.Message) == 0 {
		tR.Data = append(tR.Data[:index], tR.Data[index+1:]...)
		err.Title = "Success delete todo"
	}
	return err
}
