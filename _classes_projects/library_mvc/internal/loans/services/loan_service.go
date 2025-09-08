package services

import (
	"errors"
	"time"

	books "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/models"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/models"
	users "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/models"
)

type LoanService struct {
	loanRepository models.LoanRepository
	userService    users.UserService
	bookService    books.BookService
}

func NewLoanService(
	loanRepository models.LoanRepository,
	userService users.UserService,
	bookService books.BookService,
) *LoanService {
	return &LoanService{
		loanRepository: loanRepository,
		userService:    userService,
		bookService:    bookService,
	}
}

func (s *LoanService) CreateLoan(bookId int64, userId int64) (*models.Loan, error) {
	book, err := s.bookService.GetBook(bookId)

	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetUser(userId)

	if err != nil {
		return nil, err
	}

	userLoans, err := s.loanRepository.GetActiveUserLoans(user.ID)

	if len(userLoans) >= 0 {
		return nil, errors.New("user has active loans")
	}

	bookLoans, err := s.loanRepository.GetActiveBookLoans(book.ID)

	if len(bookLoans) >= book.Quantity {
		return nil, errors.New("no available copies for this book")
	}

	loan := &models.Loan{
		UserID:     user.ID,
		BookID:     book.ID,
		BorrowedAt: time.Now(),
		Status:     models.STATUS_ACTIVE,
	}

	return s.loanRepository.CreateLoan(loan)
}
func (s *LoanService) ReturnLoan(loanId int64) error {
	loan, err := s.loanRepository.GetLoanById(loanId)
	if err != nil {
		return errors.New("loan not found")
	}
	if loan.Status == models.STATUS_RETURNED {
		return errors.New("loan already returned")
	}

	return s.loanRepository.ReturnLoan(loanId)
}

func (s *LoanService) GetAllLoans() ([]*models.Loan, error) {
	return s.loanRepository.GetAllLoans()
}

func (s *LoanService) GetLoan(loanId int64) (*models.Loan, error) {
	return s.loanRepository.GetLoanById(loanId)
}

func (s *LoanService) GetUserLoans(userId int64) ([]*models.Loan, error) {
	return s.loanRepository.GetActiveUserLoans(userId)
}
