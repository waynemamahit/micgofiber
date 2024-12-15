package controller

import (
	"micgofiber/lib"
	"micgofiber/model"
	"micgofiber/service"
	"os"

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
		return c.SendStatus(400)
	}
	errors := lib.ValidateResponse{}
	if errors.ValidateStruct(*request) != nil {
		return c.Status(400).JSON(errors)
	}
	response := tC.TodoService.Action(request)
	if len(response.Message) > 0 {
		return c.Status(400).JSON(response)
	}

	return c.JSON(response)
}

func (tC *TodoController) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(model.TodoResponse{
			Message: "File upload failed",
		})
	}

	// Define the destination path
	destPath := "../storage/" + file.Filename
	// Ensure the directory exists
	err = os.MkdirAll("../storage", os.ModePerm)
	if err != nil {
		return c.Status(500).SendString("Failed to create directory")
	}
	// Save the file to the server
	err = c.SaveFile(file, destPath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "File save failed",
		})
	}

	resp := tC.TodoService.SaveLogFile(model.TodoFile{
		Filename:    file.Filename,
		Description: c.FormValue("description", "This is default log description."),
	})

	return c.JSON(resp)
}
