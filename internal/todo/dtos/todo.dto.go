package todoDTO

type CreateTodoRequest struct {
	Title string
}

type EditTodoRequest struct {
	CreateTodoRequest
	Done bool
}

type CreateTodoResponse struct {
	Title string `json:"title"`
	Id    string `json:"id"`
	Done  bool   `json:"status"`
}

type TodoResponse struct {
	Title string `json:"title"`
	Id    string `json:"id"`
	Done  bool   `json:"status"`
}
