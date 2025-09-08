package models

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookById(id int64) (*Book, error)
	GetBookByTitle(title string) (*Book, error)
	GetAllBooks() ([]*Book, error)
	UpdateBook(id int64, book *Book) error
	DeleteBook(id int64) error
}
