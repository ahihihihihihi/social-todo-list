package biz

import (
	"context"
	"fmt"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTiTleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data) ; err != nil {
		fmt.Println("err from biz: ",err)
		return common.ErrCanNotCreateEntity(model.EntityName,err)
	}

	return nil

}
