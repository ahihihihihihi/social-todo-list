package main

import (
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

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;`
	Title       string     `json:"title gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;`
	Status      string     `json:"status" gorm:"column:status;`
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
			items.GET("")
			items.GET("/:id",GetItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")
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

		data.Id = id

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