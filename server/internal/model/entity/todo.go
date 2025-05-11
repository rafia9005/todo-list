package entity

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"type:text;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*Todo) TableName() string {
	return "todos"
}
