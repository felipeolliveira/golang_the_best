package repositories

import (
	"errors"
	"sync"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/models"
)

type BookInMemoryRepository struct {
	books  map[int64]*models.Book
	mu     sync.RWMutex
	nextId int64
}

func NewBookInMemoryRepository() *BookInMemoryRepository {
	return &BookInMemoryRepository{
		books:  make(map[int64]*models.Book),
		mu:     sync.RWMutex{},
		nextId: 1,
	}
}

func (r *BookInMemoryRepository) CreateBook(book *models.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	book.ID = r.nextId
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	r.books[r.nextId] = book
	return nil
}

func (r *BookInMemoryRepository) GetBookById(id int64) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	book, ok := r.books[id]

	if !ok {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (r *BookInMemoryRepository) GetBookByTitle(title string) (*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, book := range r.books {
		if book.Title == title {
			return book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (r *BookInMemoryRepository) GetAllBooks() ([]*models.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	booksList := make([]*models.Book, 0, len(r.books))
	for _, book := range r.books {
		booksList = append(booksList, book)
	}

	return booksList, nil
}

func (r *BookInMemoryRepository) UpdateBook(id int64, book *models.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	updateBook, ok := r.books[id]

	if !ok {
		return errors.New("book not found")
	}

	updateBook.Title = book.Title
	updateBook.Author = book.Author
	updateBook.Quantity = book.Quantity
	updateBook.UpdatedAt = time.Now()

	r.books[id] = updateBook

	return nil
}
func (r *BookInMemoryRepository) DeleteBook(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.books, id)

	return nil
}
