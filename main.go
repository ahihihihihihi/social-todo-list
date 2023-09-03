package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [3]string{"Doing","Done","Deleted"}

func (item ItemStatus) String() string {
	return allItemStatuses[item]
}

func parseStrToItemStatus(s string) (ItemStatus,error)  {
	for i := range allItemStatuses{
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("Invalid status string!")
}

func (item *ItemStatus) Scan(value interface{}) error  {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s",value))
	}

	v, err := parseStrToItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s",value))
	}

	*item = v

	return nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error)  {
	return []byte(fmt.Sprintf("\"%s\"",item.String())), nil
}

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;`
	Title       string     `json:"title gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;`
	Status      *ItemStatus     `json:"status" gorm:"column:status;`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;`
}

func (TodoItem) TableName() string  {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int        `json:"-" gorm:"column:id;`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	//Status      string     `json:"status" gorm:"column:description;"`
}

func (TodoItemCreation) TableName() string  {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title;"`
	Description *string     `json:"description" gorm:"column:description;"`
	Status      *string     `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string  {
	return TodoItem{}.TableName()
}

type Paging struct {
	Page int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process()  {
	if p.Page <= 0 {
p.Page = 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}

func main() {
	fmt.Println("BEGIN")
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("database: ",db)
	//
	//now := time.Now().UTC()
	//
	//item := TodoItem{
	//	Id:          1,
	//	Title:       "This is item 1",
	//	Description: "This is item 1 description",
	//	Status:      "Doing",
	//	CreatedAt:   &now,
	//	UpdatedAt:   &now,
	//}

	r := gin.Default()

	//CRUD: Create, Read, Update, Delete
	//POST: /v1/items (create a new item)
	//GET: /v1/items (list items) /v1/items?page=1
	//GET: /v1/items/:id (get item by id)
	//(PUT || PATCH) /v1/items/:id (update an item by id)
	//DELETE /v1/items/:id (delete an item by id)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreatItem(db))
			items.GET("",ListItem(db))
			items.GET("/:id",GetItem(db))
			items.PATCH("/:id",UpdateItem(db))
			items.DELETE("/:id",DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func CreatItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var data TodoItemCreation

		if err := c.ShouldBind(&data) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		fmt.Println("data2: ",data)
		fmt.Println("data: ",&data)

		if err := db.Create(&data).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var data TodoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//data.Id = id

		if err := db.Where("id = ?",id).First(&data).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}

}

func UpdateItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var data TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		fmt.Println("id: ",id)


		//parse json use ShouldBindJSON, it works
		if err := c.ShouldBindJSON(&data) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//fmt.Println("data2: ",data)
		//fmt.Printf("data %v, %T: ",data,data)
		//fmt.Println("data3: ",string(*data.Title))
		res2B, _ := json.Marshal(data)
		fmt.Println(string(res2B))

		if err := db.Where("id = ?",id).Updates(&data).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}

}

func DeleteItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//if err := db.Table(TodoItem{}.TableName()).Where("id = ?",id).Delete(nil).Error ; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//
		//	return
		//}

		if err := db.Table(TodoItem{}.TableName()).Where("id = ?",id).Updates(map[string]interface{}{
			"status":"Deleted",
		}).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}

}

func ListItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {

		var paging Paging

		if err := c.ShouldBind(&paging) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		var result []TodoItem

		db = db.Where("status <> ?","Deleted")

		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).Error ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
			"paging":paging,
		})
	}

}