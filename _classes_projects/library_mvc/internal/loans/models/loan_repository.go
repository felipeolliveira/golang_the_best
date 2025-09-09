package models

type LoanRepository interface {
	GetLoanById(id int64) (*Loan, error)
	GetActiveUserLoans(userId int64) []*Loan
	GetActiveBookLoans(bookId int64) []*Loan
	GetAllLoans() []*Loan
	CreateLoan(loan *Loan) (*Loan, error)
	ReturnLoan(loanId int64) error
}
