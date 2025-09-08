package models

import "time"

type Loan struct {
	ID         int64      `json:"id"`
	BookID     int64      `json:"book_id" binding:"required"`
	UserID     int64      `json:"user_id" binding:"required"`
	Status     LoanStatus `json:"status"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt time.Time  `json:"returned_at"`
}

type LoanStatus int

const (
	STATUS_ACTIVE LoanStatus = iota
	STATUS_RETURNED
)

func (ls LoanStatus) String() string {
	switch ls {
	case STATUS_ACTIVE:
		return "ACTIVE"
	case STATUS_RETURNED:
		return "RETURNED"
	default:
		return ""
	}
}
