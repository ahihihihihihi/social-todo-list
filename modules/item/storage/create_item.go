package storage

import (
	"context"
	"fmt"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		fmt.Println("err from storage: ",err)
		return common.ErrDB(err)
	}

	return nil
}
