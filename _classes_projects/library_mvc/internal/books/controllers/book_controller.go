package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/models"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	service models.BookService
}

func NewBookController(service models.BookService) *BookController {
	return &BookController{service: service}
}

func (c *BookController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")

	books.POST("", c.CreateBook)
	books.GET("/:id", c.GetBook)
	books.GET("", c.GetAllBooks)
	books.PUT("/:id", c.UpdateBook)
	books.DELETE("/:id", c.DeleteBook)
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateBook(&book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (c *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid book ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := c.service.GetBook(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid book ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.UpdateBook(id, &book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid book ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = c.service.DeleteBook(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
