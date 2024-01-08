package repository

import (
	"micgofiber/db"
	"micgofiber/lib"
	"micgofiber/model"
)

type TodoRepo struct {
	Data []db.Todo
	Db   *lib.DBConfig
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{
		Db: lib.NewDB(),
	}
}

func (repo *TodoRepo) AddTodo(newTodo db.Todo) model.TodoResponse {
	repo.Data = append(repo.Data, newTodo)
	return model.TodoResponse{
		Title: "Success add todo",
	}
}

func (repo *TodoRepo) findTodo(index int, action string) model.TodoResponse {
	for i := range repo.Data {
		if i == index {
			return model.TodoResponse{}
		}
	}
	return model.TodoResponse{
		Title:   action + " Failed",
		Message: "Data Not Found",
	}
}

func (repo *TodoRepo) UpdateTodo(newTodo db.Todo, index int) model.TodoResponse {
	err := repo.findTodo(index, "Update")
	if len(err.Message) == 0 {
		repo.Data[index] = newTodo
		err.Title = "Success update todo"
	}
	return err
}

func (repo *TodoRepo) DeleteTodo(index int) model.TodoResponse {
	err := repo.findTodo(index, "Delete")
	if len(err.Message) == 0 {
		repo.Data = append(repo.Data[:index], repo.Data[index+1:]...)
		err.Title = "Success delete todo"
	}
	return err
}
