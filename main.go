package main

import (
	"micgofiber/lib"
	"micgofiber/router"
)

func main() {
	server := lib.NewApp()

	router.NewTodoRouter(server)

	server.App.Listen(":" + server.Port)
}
