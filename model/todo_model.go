package model

type TodoModel struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Check       bool   `json:"check"`
}

type TodoRequest struct {
	Action string    `json:"action" validate:"required"`
	Data   TodoModel `json:"data" validate:"required"`
	Index  int       `json:"index" validate:"min=0"`
}

type TodoResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}
