package router

import (
	"micgofiber/controller"
	"micgofiber/lib"
	"micgofiber/repository"
	"micgofiber/service"
)

func NewTodoRouter(app *lib.AppConfig) {
	router := app.GetRouterV1("/todo")

	tR := repository.NewTodoRepo()
	tS := service.NewTodoService(tR)
	tC := controller.TodoController{TodoService: tS}

	router.Get("/", tC.GetTodo)
	router.Post("/", tC.ActionTodo)
}
