package test

import (
	"micgofiber/lib"
	"micgofiber/test/mock"
	"net/http"
)

type TodoE2E struct {
	app  *lib.AppConfig
	csrf *CsrfE2E
	mock *mock.TodoMock
}

func NewTodoE2E(app *lib.AppConfig) *TodoE2E {
	return &TodoE2E{
		app,
		NewCsrfE2E(app),
		mock.NewTodoMock(),
	}
}

func (e2e *TodoE2E) GetTodo() (*http.Response, error) {
	return e2e.csrf.Request("/todo", "GET", nil)
}

func (e2e *TodoE2E) ActionTodo() (*http.Response, error) {
	return e2e.csrf.Request("/todo", "POST", e2e.mock.Dto)
}

func (e2e *TodoE2E) UploadFile() (*http.Response, error) {
	return e2e.csrf.FormRequest("/todo", "PUT", "file", "test_mock.txt", e2e.mock.LogFile)
}
