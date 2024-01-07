package mock

import (
	"micgofiber/model"

	"github.com/jaswdr/faker"
)

type TodoMock struct {
	Dto *model.TodoRequest
}

func NewTodoMock() *TodoMock {
	fake := faker.New()

	return &TodoMock{
		Dto: &model.TodoRequest{
			Action: "add",
			Data: model.TodoModel{
				Title:       "New Todo",
				Description: fake.Lorem().Paragraph(200),
				Check:       false,
			},
			Index: 1,
		},
	}
}
