package controller

import (
	"micgofiber/lib"
	"micgofiber/model"
	"micgofiber/service"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	TodoService *service.TodoService
}

func (tC *TodoController) GetTodo(c *fiber.Ctx) error {
	response := tC.TodoService.GetData()
	return c.JSON(response)
}

func (tC *TodoController) ActionTodo(c *fiber.Ctx) error {
	request := new(model.TodoRequest)
	if err := c.BodyParser(request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	errors := lib.ValidateResponse{}
	if errors.ValidateStruct(*request) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	response := tC.TodoService.Action(request)
	if len(response.Message) > 0 {
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return c.JSON(response)
}
