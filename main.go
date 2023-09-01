package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	fmt.Println(`Hello golang`)

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "This is item 1",
		Description: "This is item 1 description",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsonData, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("item: ", string(jsonData))
	//fmt.Println(jsonData)

	jsonStr := "{\"id\":1,\"title\":\"This is item 1\",\"description\":\"This is item 1 description\",\"status\":\"Doing\",\"created_at\":\"2023-09-01T07:45:34.6859533Z\"}"

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("item2: ", item2)

}
