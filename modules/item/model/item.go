package model

import (
	"errors"
	"social-todo-list/common"
)

const (
	EntityName = "Item"
)

var (
	ErrTiTleIsBlank = errors.New("Title can not be blank!")
	ErrItemDeleted = errors.New("Item is deleted!")

	ErrItemDeletedNew = common.NewCustomError(errors.New("Item is deleted!"),"Item has been deleted!","ErrItemDeleted")
)

type TodoItem struct {
	common.SQLMODEL
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
