package main

import (
	"log"

	booksControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/books/controllers"
	loansControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/loans/controllers"
	usersControllers "github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_ = router

	usersControllers := usersControllers.NewUserController()
	booksControllers := booksControllers.NewBookController()
	loansControllers := loansControllers.NewLoanController()

	booksControllers.RegisterRoutes(router)
	loansControllers.RegisterRoutes(router)
	usersControllers.RegisterRoutes(router)

	if err := router.Run(":3130"); err != nil {
		log.Fatal(err)
	}
}
