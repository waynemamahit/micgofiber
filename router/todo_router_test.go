package router

import (
	"micgofiber/lib"
	"micgofiber/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoRouter(t *testing.T) {
	app := lib.NewApp()
	NewTodoRouter(app)
	todo := test.NewTodoE2E(app)
	errMessage := "Failed to make request!"

	{
		resp, err := todo.GetTodo()
		if err != nil {
			assert.Error(t, err, errMessage)
		} else {
			assert.Equalf(t, 200, resp.StatusCode, "should get todo")
		}
	}

	{
		resp, err := todo.ActionTodo()
		if err != nil {
			assert.Error(t, err, errMessage)
		} else {
			assert.Equalf(t, 200, resp.StatusCode, "should add new todo")
		}
	}

	{
		resp, err := todo.UploadFile()
		if err != nil {
			assert.Error(t, err, errMessage)
		} else {
			assert.Equalf(t, 200, resp.StatusCode, "should upload a file")
		}
	}

}
