package http

import (
	"net/http"
	"researching-go/rest_api/todo"
)

type HTTPHandlers struct {
	todolist *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{todolist: todoList}
}

/*
info:    JSON in HTTP request body
method:  POST
pattern: /task

succeed:
  - status code: 201 Created
  - response body: JSON represent created task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}
