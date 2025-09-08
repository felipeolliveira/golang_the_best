package models

import "time"

type Book struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required,min=5"`
	Author    string    `json:"author" binding:"required,min=5"`
	Quantity  int       `json:"quantity" binding:"required,gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
