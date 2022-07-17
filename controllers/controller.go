package controllers

import (
	"net/http"

	"github.com/ashishsingh4u/bookmicroservice/models"
	"github.com/ashishsingh4u/bookmicroservice/repository"
	"github.com/gin-gonic/gin"
)

var repo repository.BookRepoInterface = &repository.BookRepository{}

// @BasePath /v1
// GetBooks godoc
// @Summary Get details of all books
// @Schemes
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(ctx *gin.Context) {
	var books []models.Book

	if err := repo.GetBooks(&books); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

// @BasePath /v1
// CreateBook godoc
// @Summary Creates book entry
// @Schemes
// @Description Creates book entry
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.CreateBookInput true "Create book"
// @Success 201 {} models.Book
// @Router /books [post]
func CreateBook(ctx *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	var book models.Book
	if err := repo.CreateBook(&input, &book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"book": book})
}

// @BasePath /v1
// FindBook godoc
// @Summary Find book
// @Schemes
// @Description Find book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the book"
// @Success 200 {} models.Book
// @Router /books/{id} [get]
func FindBook(ctx *gin.Context) { // Get model if exist
	var book models.Book

	bookId := ctx.Param("id")
	if err := repo.GetBook(bookId, &book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

// @BasePath /v1
// UpdateBook godoc
// @Summary Update a book
// @Schemes
// @Description Update a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the book"
// @Param book body models.UpdateBookInput true "Update book"
// @Success 200 {} models.Book
// @Router /books/{id} [patch]
func UpdateBook(ctx *gin.Context) {
	// Validate input
	var input models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if code, err := repo.UpdateBook(ctx.Param("id"), &input, &book); err != nil {
		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

// @BasePath /v1
// DeleteBook godoc
// @Summary Delete book
// @Schemes
// @Description Delete book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the book"
// @Success 200
// @Router /books/{id} [delete]
func DeleteBook(ctx *gin.Context) {

	if err := repo.DeleteBook(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "book deleted"})
}
