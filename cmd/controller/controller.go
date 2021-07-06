package controller

import (
	"encoding/json"
	"net/http"
)

type Controller struct{}

func NewController(s *http.Server) *Controller {
	c := &Controller{}

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
	}
}
