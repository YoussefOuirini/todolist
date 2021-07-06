package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ToDo struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	DueDate     time.Time      `json:"dueDate"`
	Labels      pq.StringArray `gorm:"type:text[]" json:"labels"`
	IsDone      bool           `json:"isDone"`
}
