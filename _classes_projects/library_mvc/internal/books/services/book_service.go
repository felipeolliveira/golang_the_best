package services

import (
	"errors"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/models"
)

type BookService struct {
	bookRepository models.BookRepository
}

func NewBookService(bookRepository models.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}

func (s *BookService) CreateBook(book *models.Book) error {
	foundBook, _ := s.bookRepository.GetBookByTitle(book.Title)

	if foundBook != nil {
		return errors.New("book with the same title already exists")
	}

	return s.bookRepository.CreateBook(book)
}

func (s *BookService) GetBook(id int64) (*models.Book, error) {
	return s.bookRepository.GetBookById(id)
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	return s.bookRepository.GetAllBooks()
}

func (s *BookService) UpdateBook(id int64, book *models.Book) error {
	return s.bookRepository.UpdateBook(id, book)
}

func (s *BookService) DeleteBook(id int64) error {
	return s.bookRepository.DeleteBook(id)
}
