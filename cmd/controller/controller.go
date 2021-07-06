package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
	http.HandleFunc("/todo/", c.getToDo)

	return c
}

func (c *Controller) handleToDo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.createToDo(w, r)
	case http.MethodGet:
		// c.getToDos(w, r)
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

		return
	}
}

func (c Controller) getToDo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todo/"):]
	toDoID, err := uuid.Parse(id)
	if err != nil || toDoID == uuid.Nil {
		http.Error(w, fmt.Sprintf("invalid uuid: %s", err.Error()), http.StatusBadRequest)

		return
	}

	toDo, err := c.toDoRepository.GetToDo(c.db, toDoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting todo: %s", err.Error()), http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(toDo)
}
