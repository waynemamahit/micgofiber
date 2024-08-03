package router

import (
	"micgofiber/controller"
	"micgofiber/lib"
)

func NewTodoRouter(app *lib.AppConfig, tC controller.TodoController) {
	router := app.GetRouterV1("/todo")

	router.Get("/", tC.GetTodo)
	router.Post("/", tC.ActionTodo)
	router.Put("/", tC.UploadFile)
}
