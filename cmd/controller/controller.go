package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/youssefouirini/todolist/cmd/model"
	"github.com/youssefouirini/todolist/cmd/storage"
)

type Controller struct {
	db             *gorm.DB
	toDoRepository storage.ToDoRepository
}

func NewController(s *http.Server, db *gorm.DB, toDoRepository storage.ToDoRepository) *Controller {
	c := &Controller{
		db:             db,
		toDoRepository: toDoRepository,
	}

	http.HandleFunc("/todo", c.handleToDo)

	return c
}

func (c *Controller) handleToDo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.createToDo(w, r)
	default:
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func (c Controller) createToDo(w http.ResponseWriter, r *http.Request) {
	var request createRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = request.validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	todo := model.ToDo{
		Title:       request.Title,
		Description: request.Description,
		Labels:      request.Labels,
		DueDate:     request.DueDate,
		IsDone:      false,
	}

	err = c.toDoRepository.CreateToDo(c.db, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
