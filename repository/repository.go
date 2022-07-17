package repository

import "github.com/ashishsingh4u/bookmicroservice/models"

type BookRepoInterface interface {
	GetBook(bookId string, book *models.Book) (err error)
	DeleteBook(bookId string) (err error)
	GetBooks(book *[]models.Book) (err error)
	CreateBook(bookInput *models.CreateBookInput, book *models.Book) (err error)
	UpdateBook(bookId string, bookInput *models.UpdateBookInput, book *models.Book) (code int, err error)
}
