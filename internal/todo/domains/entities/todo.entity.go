package todoEntity

import (
	"github.com/oklog/ulid/v2"
	todoDTO "github.com/zombieleet/codesphere-test-todo-app/internal/todo/dtos"
)

// public interace for persistence needs
type TodoSnapshot interface {
	ID() string
	Title() string
	IsDone() bool
}

type todoEntity struct {
	title string
	done  bool
	id    string
}

func NewTodo(todo todoDTO.CreateTodoRequest) todoEntity {
	return todoEntity{
		id:    ulid.Make().String(),
		title: todo.Title,
		done:  false,
	}
}

func HydrateTodo(id string, title string, done bool) TodoSnapshot {
	return todoEntity{
		id:    id,
		title: title,
		done:  done,
	}
}

func (t todoEntity) EditTodo(opt todoDTO.EditTodoRequest) todoEntity {
	t.title = opt.Title
	t.done = opt.Done
	return t
}

func (t todoEntity) Equals(todo todoEntity) bool {
	return t.id == todo.id
}

func (t todoEntity) ID() string {
	return t.id
}

func (t todoEntity) Title() string {
	return t.title
}

func (t todoEntity) IsDone() bool {
	return t.done
}
