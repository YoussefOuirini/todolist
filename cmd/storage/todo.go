package storage

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/youssefouirini/todolist/cmd/model"
)

type ToDoRepository interface {
	CreateToDo(db *gorm.DB, toDo *model.ToDo) error
	GetToDo(db *gorm.DB, id uuid.UUID) (*model.ToDo, error)
}

type toDoRepository struct{}

func NewToDoRepository() ToDoRepository {
	return &toDoRepository{}
}

func (repo *toDoRepository) CreateToDo(db *gorm.DB, toDo *model.ToDo) error {
	return db.Create(toDo).Error
}

func (repo *toDoRepository) GetToDo(db *gorm.DB, id uuid.UUID) (*model.ToDo, error) {
	toDo := &model.ToDo{}

	err := db.Model(toDo).Where("id = ?", id).Take(&toDo).Error

	return toDo, err
}
