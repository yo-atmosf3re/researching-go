package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"researching-go/pkg/logger"
	"researching-go/rest_api/todo"
	"time"
)

type HTTPHandlers struct {
	todolist *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{todolist: todoList}
}

/*
pattern: /task
method:  POST
info:    JSON in HTTP request body

succeed:
  - status code: 201 Created
  - response body: JSON represent created task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todolist.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExist) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(h.todolist, "", "  ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		logger.Ptc("failed to write http response, ", err)
	}
}

/*
pattern: /tasks/{title}
method: GET
info: pattern

succeed:
  - status code: 200 Ok
  - response body: JSON represented found task

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks
method: GET
info: -

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?completed=true
method: GET
info: query params

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: PATCH
info: pattern + JSON in request body

succeed:
  - status code: 200 Ok
  - response body: JSON represented changed task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
  - status code: 204 No Content
  - response body: -

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
