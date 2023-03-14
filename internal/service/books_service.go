package service

import (
	"context"
	"log"

	"github.com/sog01/simple-bookstore/internal/model"
	"github.com/sog01/simple-bookstore/internal/repository"
)

// Books represent books service contract
type Books interface {
	GetBooks(ctx context.Context, req *model.GetBooksRequest) (*model.GetBooksResponse, error)
}

// BooksService represent books service object implementation
type BooksService struct {
	repo repository.Books
}

// NewBooksService construct new books service object
func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{
		repo: repo,
	}
}

// GetBooks get books service
func (s *BooksService) GetBooks(ctx context.Context, req *model.GetBooksRequest) (*model.GetBooksResponse, error) {
	if req.Page <= 1 {
		req.Page = 0
	}
	from := req.Page * req.Size
	bookList, total, err := s.repo.GetBooks(ctx, from, req.Size)
	if err != nil {
		log.Printf("failed get books: %v\n", err)
		return nil, err
	}
	return &model.GetBooksResponse{
		BookList: bookList,
		Total:    total,
		MaxPage:  total / req.Size,
	}, nil
}
