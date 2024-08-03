package router

import (
	"micgofiber/controller"
	"micgofiber/lib"
	"micgofiber/repository"
	"micgofiber/service"
)

func InitApp(app *lib.AppConfig) {
	tR := repository.NewTodoRepo()
	tS := service.NewTodoService(tR)

	NewTodoRouter(app, controller.TodoController{TodoService: tS})
}
