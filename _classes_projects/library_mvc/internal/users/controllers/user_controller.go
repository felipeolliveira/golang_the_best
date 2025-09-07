package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service models.UserService
}

func NewUserController(service models.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("", c.CreateUser)
	users.GET("/:id", c.GetUser)
	users.GET("", c.GetAllUser)
	users.PUT("/:id", c.UpdateUser)
	users.DELETE("/:id", c.DeleteUser)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.service.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid user ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.service.GetUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid user ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = c.service.UpdateUser(id, &user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		slog.Error("Invalid user ID", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.service.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
