package model

import "time"

type ToDo struct {
	Title       string
	Description string
	DueDate     time.Time
	Labels      []string
	IsDone      bool
}
