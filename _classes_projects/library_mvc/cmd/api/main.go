package main

import (
	"log"

	bookControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/controllers"
	bookRepositories "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/repositories"
	bookServices "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/services"

	loanControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/controllers"
	loanRepositories "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/repositories"
	loanServices "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/services"

	userControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/controllers"
	userRepositories "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/repositories"
	userServices "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_ = router

	userRepository := userRepositories.NewUserInMemoryRepository()
	userService := userServices.NewUserService(userRepository)
	userController := userControllers.NewUserController(userService)

	bookRepository := bookRepositories.NewBookInMemoryRepository()
	bookService := bookServices.NewBookService(bookRepository)
	bookController := bookControllers.NewBookController(bookService)

	loanRepository := loanRepositories.NewLoanInMemoryRepository()
	loanService := loanServices.NewLoanService(loanRepository, userService, bookService)
	loansController := loanControllers.NewLoanController(loanService)

	bookController.RegisterRoutes(router)
	loansController.RegisterRoutes(router)
	userController.RegisterRoutes(router)

	if err := router.Run(":3130"); err != nil {
		log.Fatal(err)
	}
}
