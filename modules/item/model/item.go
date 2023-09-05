package model

import (
	"errors"
	"social-todo-list/common"
)

var (
	ErrTiTleIsBlank = errors.New("Title can not be blank!")
)

type TodoItem struct {
	common.SQLMODEL
	Title       string      `json:"title gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;`
	Status      *ItemStatus `json:"status" gorm:"column:status;`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id;`
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
