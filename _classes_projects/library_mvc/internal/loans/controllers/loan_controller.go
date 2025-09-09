package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/models"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	service models.LoanService
}

func NewLoanController(service models.LoanService) *LoanController {
	return &LoanController{service: service}
}

func (c *LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")
	loans.POST("", c.CreateLoan)
	loans.GET("/:id", c.GetLoan)
	loans.GET("", c.GetAllLoans)
	loans.PATCH("/:id/return", c.ReturnLoan)

	users := r.Group("/loans/users")
	users.GET("/:userId", c.GetUserLoans)
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookId int64 `json:"book_id"`
		UserId int64 `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := c.service.CreateLoan(request.BookId, request.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (c *LoanController) GetLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid loan ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := c.service.GetLoan(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {
	loans := c.service.GetAllLoans()

	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) GetUserLoans(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)

	if err != nil {
		slog.Error("Invalid userId", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loans := c.service.GetUserLoans(userId)

	ctx.JSON(http.StatusOK, loans)
}

func (c *LoanController) ReturnLoan(ctx *gin.Context) {
	loanId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid loanId", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	err = c.service.ReturnLoan(loanId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
