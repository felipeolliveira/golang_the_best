package models

type LoanService interface {
	CreateLoan(bookId int64, userId int64) (*Loan, error)
	GetLoan(loanId int64) (*Loan, error)
	GetAllLoans() ([]*Loan, error)
	GetUserLoans(userId int64) ([]*Loan, error)
	ReturnLoan(loanId int64) error
}
