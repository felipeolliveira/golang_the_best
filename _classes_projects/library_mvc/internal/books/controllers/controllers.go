package books

import "github.com/gin-gonic/gin"

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

func (c *BookController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")

	books.POST("", c.CreateBook)
	books.GET("/:id", c.GetBook)
	books.GET("", c.GetAllBooks)
	books.PUT("/:id", c.UpdateBook)
	books.DELETE("/:id", c.DeleteBook)
}

func (c *BookController) CreateBook(ctx *gin.Context) {}

func (c *BookController) GetBook(ctx *gin.Context) {}

func (c *BookController) GetAllBooks(ctx *gin.Context) {}

func (c *BookController) UpdateBook(ctx *gin.Context) {}

func (c *BookController) DeleteBook(ctx *gin.Context) {}
