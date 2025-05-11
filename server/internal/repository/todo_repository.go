package repository

import (
	"todo-list/internal/model/entity"
	"todo-list/pkg/database"
)

func GetAll() (*entity.Todo, error) {
	var todos *entity.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func FindById(id string) (*entity.Todo, error) {
	var todos *entity.Todo
	if err := database.DB.First(&todos, id).Error; err != nil {
		return nil, err
	}

	return todos, nil
}
