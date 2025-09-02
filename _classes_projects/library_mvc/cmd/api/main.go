package main

import (
	"log"

	books "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/controllers"
	loans "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/controllers"
	users "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_ = router

	usersControllers := users.NewUserController()
	booksControllers := books.NewBookController()
	loansControllers := loans.NewLoanController()

	booksControllers.RegisterRoutes(router)
	loansControllers.RegisterRoutes(router)
	usersControllers.RegisterRoutes(router)

	if err := router.Run(":3130"); err != nil {
		log.Fatal(err)
	}
}
