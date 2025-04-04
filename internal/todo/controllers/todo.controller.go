package todoController

import (
	"encoding/json"
	"fmt"
	"net/http"

	todoDTO "github.com/zombieleet/codesphere-test-todo-app/internal/todo/dtos"
	todoService "github.com/zombieleet/codesphere-test-todo-app/internal/todo/services"
)

type todoController struct {
	todoSvc todoService.TodoService
}

func InitTodoController() *http.ServeMux {
	todoMux := http.NewServeMux()
	todoController := todoController{todoSvc: todoService.NewTodoService()}
	todoMux.HandleFunc("POST /{$}", todoController.createTodo)
	todoMux.HandleFunc("GET /{id}", todoController.getTodo)
	return todoMux
}

func (controller todoController) createTodo(rw http.ResponseWriter, rq *http.Request) {

	body := rq.Context().Value(struct{}{}).(map[string]interface{})

	title, ok := body["title"].(string)

	if !ok {
		rw.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprintln(rw, "Cannot read/parse required todo Title")
		return
	}

	todo, err := controller.todoSvc.CreateTodo(todoDTO.CreateTodoRequest{
		Title: title,
	})

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
		return
	}

	todoAsByte, err := json.Marshal(todo)

	if err != nil {
		rw.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprintln(rw, err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(todoAsByte)
}

func (controller todoController) getTodo(rw http.ResponseWriter, rq *http.Request) {
	todoId := rq.PathValue("id")

	todo, err := controller.todoSvc.GetTodo(todoId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
		return
	}

	todoAsByte, err := json.Marshal(todo)

	if err != nil {
		rw.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprintln(rw, err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(todoAsByte)
}
