package ginitem

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		fmt.Println("id: ", id)

		//parse json use ShouldBindJSON, it works
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		//fmt.Println("data2: ",data)
		//fmt.Printf("data %v, %T: ",data,data)
		//fmt.Println("data3: ",string(*data.Title))
		res2B, _ := json.Marshal(data)
		fmt.Println(string(res2B))

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if	err := business.UpdateItemById(c.Request.Context(),id,&data)	; err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
