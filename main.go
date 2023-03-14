package main

import (
	"github.com/sog01/simple-bookstore/internal/handler"
	"github.com/sog01/simple-bookstore/internal/repository"
	"github.com/sog01/simple-bookstore/internal/service"
)

func main() {
	booksRepo := repository.NewBooksRepository("http://localhost:9200")
	booksService := service.NewBooksService(booksRepo)

	rest := handler.NewRest(booksService)
	rest.ListenAndServe()
}
