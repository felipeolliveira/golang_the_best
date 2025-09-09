package repositories

import (
	"errors"
	"sync"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/models"
)

type LoanInMemoryRepository struct {
	loans  map[int64]*models.Loan
	mu     sync.RWMutex
	nextId int64
}

func NewLoanInMemoryRepository() *LoanInMemoryRepository {
	return &LoanInMemoryRepository{
		loans:  make(map[int64]*models.Loan),
		mu:     sync.RWMutex{},
		nextId: 1,
	}
}

func (r *LoanInMemoryRepository) GetLoanById(id int64) (*models.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	loan, ok := r.loans[id]
	if !ok {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

func (r *LoanInMemoryRepository) GetActiveUserLoans(userId int64) []*models.Loan {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.loans) == 0 {
		return []*models.Loan{}
	}

	activeUserLoans := make([]*models.Loan, 0)
	for loanId, loan := range r.loans {
		if loan.UserID == userId && loan.Status == models.STATUS_ACTIVE {
			activeUserLoans = append(activeUserLoans, r.loans[loanId])
		}
	}

	return activeUserLoans
}

func (r *LoanInMemoryRepository) GetActiveBookLoans(bookId int64) []*models.Loan {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.loans) == 0 {
		return []*models.Loan{}
	}

	activeBookLoans := make([]*models.Loan, 0)
	for loanId, loan := range r.loans {
		if loan.BookID == bookId && loan.Status == models.STATUS_ACTIVE {
			activeBookLoans = append(activeBookLoans, r.loans[loanId])
		}
	}

	return activeBookLoans
}

func (r *LoanInMemoryRepository) GetAllLoans() []*models.Loan {
	r.mu.Lock()
	defer r.mu.Unlock()

	loansList := make([]*models.Loan, 0, len(r.loans))

	for _, loan := range r.loans {
		loansList = append(loansList, loan)
	}

	return loansList
}

func (r *LoanInMemoryRepository) CreateLoan(loan *models.Loan) (*models.Loan, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	loan.ID = r.nextId
	loan.Status = models.STATUS_ACTIVE
	loan.BorrowedAt = time.Now()

	r.loans[r.nextId] = loan
	r.nextId++

	return loan, nil
}

func (r *LoanInMemoryRepository) ReturnLoan(loanId int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	loan, ok := r.loans[loanId]

	if !ok {
		return errors.New("loan not found")
	}

	loan.Status = models.STATUS_RETURNED
	loan.ReturnedAt = time.Now()

	r.loans[loanId] = loan

	return nil
}
