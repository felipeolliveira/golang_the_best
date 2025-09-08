package models

type LoanRepository interface {
	GetLoanById(id int64) (*Loan, error)
	GetActiveUserLoans(userId int64) ([]*Loan, error)
	GetActiveBookLoans(bookId int64) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
	CreateLoan(loan *Loan) (*Loan, error)
	ReturnLoan(loanId int64) error
}
