package loans

import "github.com/gin-gonic/gin"

type LoanController struct {
}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (c *LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")

	loans.POST("", c.CreateLoan)
	loans.GET("/:id", c.GetLoan)
	loans.GET("", c.GetAllLoans)
	loans.PUT("/:id", c.UpdateLoan)
	loans.DELETE("/:id", c.DeleteLoan)
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {}

func (c *LoanController) GetLoan(ctx *gin.Context) {}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {}
