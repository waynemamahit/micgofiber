package model

import "micgofiber/db"

type TodoRequest struct {
	Action string  `json:"action" validate:"required"`
	Data   db.Todo `json:"data" validate:"required"`
	Index  int     `json:"index" validate:"min=0"`
}

type TodoResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type TodoFile struct {
	Filename    string `json:"filename"`
	Description string `json:"description"`
}
