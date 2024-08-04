package router

import (
	"micgofiber/controller"
	"micgofiber/lib"
	"micgofiber/repository"
	"micgofiber/service"
)

func InitApp(app *lib.AppConfig) {
	db := lib.NewDB()
	tR := repository.NewTodoRepo(db)
	tS := service.NewTodoService(tR)

	NewTodoRouter(app, controller.TodoController{TodoService: tS})
}
