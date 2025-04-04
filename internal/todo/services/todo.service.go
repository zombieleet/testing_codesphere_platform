package todoService

import (
	"context"

	todoEntity "github.com/zombieleet/codesphere-test-todo-app/internal/todo/domains/entities"
	todoDTO "github.com/zombieleet/codesphere-test-todo-app/internal/todo/dtos"
	todoRepository "github.com/zombieleet/codesphere-test-todo-app/internal/todo/repository"
)

type TodoService struct {
	todoRepo *todoRepository.TodoRepository
}

func NewTodoService() TodoService {
	return TodoService{todoRepo: todoRepository.GetTodoRepository()}
}

func (svc TodoService) CreateTodo(createTodo todoDTO.CreateTodoRequest) (*todoDTO.CreateTodoResponse, error) {
	todo := todoEntity.NewTodo(createTodo)

	err := svc.todoRepo.Save(context.Background(), todo)

	if err != nil {
		return nil, err
	}

	todoResponse := todoDTO.CreateTodoResponse{
		Title: todo.Title(),
		Done:  todo.IsDone(),
		Id:    todo.ID(),
	}

	return &todoResponse, err

}

func (svc TodoService) GetTodo(todoId string) (*todoDTO.TodoResponse, error) {
	result, err := svc.todoRepo.GetTodo(context.Background(), todoId)

	if err != nil {
		return nil, err
	}

	todoResponse := todoDTO.TodoResponse{
		Title: result.Title(),
		Done:  result.IsDone(),
		Id:    result.ID(),
	}

	return &todoResponse, nil

}
