package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ToDo struct {
	ID          uuid.UUID
	Title       string
	Description string
	DueDate     time.Time
	Labels      pq.StringArray `gorm:"type:text[]"`
	IsDone      bool
}
