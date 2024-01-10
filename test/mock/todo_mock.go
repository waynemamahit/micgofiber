package mock

import (
	"micgofiber/db"
	"micgofiber/model"

	"github.com/jaswdr/faker"
)

type TodoMock struct {
	Dto     *model.TodoRequest
	LogFile *model.TodoFile
}

func NewTodoMock() *TodoMock {
	fake := faker.New()

	return &TodoMock{
		Dto: &model.TodoRequest{
			Action: "add",
			Data: db.Todo{
				Title:       "New Todo",
				Description: fake.Lorem().Paragraph(200),
				Check:       false,
			},
			Index: 1,
		},
		LogFile: &model.TodoFile{
			Filename:    "test_result.txt",
			Description: "Lorem ipsum",
		},
	}
}
