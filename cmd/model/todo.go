package model

import (
	"time"

	"github.com/lib/pq"
)

type ToDo struct {
	Title       string
	Description string
	DueDate     time.Time
	Labels      pq.StringArray `gorm:"type:text[]"`
	IsDone      bool
}
