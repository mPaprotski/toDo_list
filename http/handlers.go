package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"todoapp/todo"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP request body

succeed:
  - status code: 201 created
  - response body: JSON represent created task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)

	if err := h.todoList.AddTask(todoTask); err != nil {
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			writeError(w, http.StatusConflict, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	writeJSON(w, http.StatusCreated, todoTask)
}

/*
pattern: /tasks/{title}
method: GET
info: pattern

succeed:
  - status code: 200 ok
  - response body: JSON represent found task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
			return
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	writeJSON(w, http.StatusOK, task)
}

/*
pattern: /tasks
method: GET
info: pattern

succeed:
  - status code: 200 ok
  - response body: JSON represent found task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()

	writeJSON(w, http.StatusOK, tasks)
}

/*
pattern: /tasks?completed=true
method: GET
info: query params

succeed:
  - status code: 200 ok
  - response body: JSON represent found task

failed:
  - status code: 400, 500, .CTask(..
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUnCompletedTasks()

	writeJSON(w, http.StatusOK, uncompletedTasks)
}

/*
pattern: /tasks/{title}
method: PATCH
info: pattern + JSON in request body

succeed:
  - status code: 200 ok
  - response body: JSON represent changed task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if completeDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)
	} else {
		changedTask, err = h.todoList.UncompleteTask(title)

	}

	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	writeJSON(w, http.StatusOK, changedTask)
}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
  - status code: 204 no content
  - response body: -

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
