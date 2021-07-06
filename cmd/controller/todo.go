package controller

import (
	"fmt"
	"net/http"
)

type Controller struct{}

func NewController(s *http.Server) *Controller {
	c := &Controller{}

	http.HandleFunc("/todo", c.CreateToDo)

	return c
}

func (c *Controller) CreateToDo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pizza!")
}
