package repository

import (
	"errors"
	"net/http"

	"github.com/ashishsingh4u/bookmicroservice/models"
)

type BookRepository struct {
}

func (repo *BookRepository) GetBook(bookId string, book *models.Book) (err error) {
	if err = models.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		err = errors.New("database error: record not found")
	}

	return
}

func (repo *BookRepository) DeleteBook(bookId string) (err error) {

	var book models.Book
	if err = models.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		err = errors.New("database error: record not found")
		return
	}

	err = models.DB.Delete(&book).Error

	return
}

func (repo *BookRepository) GetBooks(books *[]models.Book) (err error) {
	if err = models.DB.Find(&books).Error; err != nil {
		err = errors.New("database error: record not found")
	}

	return
}

func (repo *BookRepository) CreateBook(bookInput *models.CreateBookInput, book *models.Book) (err error) {
	*book = models.Book{Title: bookInput.Title, Author: bookInput.Author}
	err = models.DB.Create(&book).Error
	return
}

func (repo *BookRepository) UpdateBook(bookId string, bookInput *models.UpdateBookInput, book *models.Book) (code int, err error) {

	if err = models.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		code = http.StatusBadRequest
		return
	}

	if err = models.DB.Model(&book).Updates(&models.Book{Title: bookInput.Title, Author: bookInput.Author}).Error; err != nil {
		err = errors.New("database error: existing record not found")
		code = http.StatusInternalServerError
	}

	return
}
