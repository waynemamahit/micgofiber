package router

import (
	"micgofiber/controller"
	"micgofiber/repository"
	"micgofiber/service"

	"github.com/gofiber/fiber/v2"
)

func NewTodoRouter(router fiber.Router) {
	tR := repository.NewTodoRepo()
	tS := service.NewTodoService(tR)
	tC := controller.TodoController{TodoService: tS}

	router.Get("/", tC.GetTodo)
	router.Post("/", tC.ActionTodo)
}
