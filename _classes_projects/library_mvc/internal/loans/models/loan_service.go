package models

type LoanService interface {
	CreateLoan(bookId int64, userId int64) (*Loan, error)
	GetLoan(loanId int64) (*Loan, error)
	GetAllLoans() []*Loan
	GetUserLoans(userId int64) []*Loan
	ReturnLoan(loanId int64) error
}
