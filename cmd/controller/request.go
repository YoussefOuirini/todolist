package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	descriptionMaxLength = 200
	labelMaxLength       = 10
)

type createRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"` //RFC 3339 format: 2021-07-12T07:20:50.52Z
	Labels      []string  `json:"labels"`
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

	if len(r.Description) > descriptionMaxLength {
		return errors.New(fmt.Sprintf("description length is %v, which is longer than max of %v", len(r.Description), descriptionMaxLength))
	}

	for _, label := range r.Labels {
		if len(label) > labelMaxLength {
			return errors.New(fmt.Sprintf("label length is %v, which is longer than max of %v", len(label), labelMaxLength))
		}
	}

	if r.ID != "" {
		if _, err := uuid.Parse(r.ID); err != nil {
			return err
		}
	}

	return nil
}
