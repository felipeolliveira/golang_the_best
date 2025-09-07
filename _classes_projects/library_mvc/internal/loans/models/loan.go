package models

import "time"

type Loan struct {
	ID         int64     `json:"id"`
	BookID     int64     `json:"book_id"`
	UserID     int64     `json:"user_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
