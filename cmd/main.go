package main

import (
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/repository"
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/service"
	"github.com/holizDV/dts-go-chapter2-challenge2/app/interfaces"
	"github.com/holizDV/dts-go-chapter2-challenge2/pkg/config"
)

func main() {

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository)
	bookController := interfaces.NewBookController(bookService)

	router := config.NewRouter(bookController)
	router.Run(":8080")
}
