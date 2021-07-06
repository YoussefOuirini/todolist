package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/youssefouirini/todolist/cmd/model"
)

type ToDoRepository interface {
	CreateToDo(db *gorm.DB, toDo *model.ToDo) error
}

type toDoRepository struct{}

func NewToDoRepository() ToDoRepository {
	return &toDoRepository{}
}

func (repo *toDoRepository) CreateToDo(db *gorm.DB, toDo *model.ToDo) error {
	return nil
}
