package common

import "time"

type SQLMODEL struct {
	Id          int         `json:"id" gorm:"column:id;`
	CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at;`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at;`
}
