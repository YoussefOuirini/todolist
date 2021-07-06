package controller

import (
	"errors"
	"time"
)

type createRequest struct {
	Title   string    `json:"title"`
	DueDate time.Time `json:"dueDate"` //RFC 3339 format: 2021-07-12T07:20:50.52Z
	Labels  []string  `json:"labels"`
}

func (r createRequest) validate() error {
	if r.Title == "" {
		return errors.New("empty title")
	}

	if r.DueDate.IsZero() {
		return errors.New("no due date")
	}

	if r.DueDate.Before(time.Now()) {
		return errors.New("due date in the past")
	}

	return nil
}
