package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
)

func ListItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		paging.Process()

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewListItemBiz(store)

		result, err := business.ListItem(c.Request.Context(),&filter,&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}