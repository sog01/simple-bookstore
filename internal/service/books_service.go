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
	GetInfiniteBooks(ctx context.Context, req *model.GetInfiniteBooksRequest) (*model.GetInfiniteBooksResponse, error)
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
	page := req.Page
	if page <= 1 {
		page = 0
	}

	size := req.Size
	if req.Size <= 1 {
		size = 6
	}
	from := page * size
	bookList, total, err := s.repo.GetBooks(ctx, from, size)
	if err != nil {
		log.Printf("failed get books: %v\n", err)
		return nil, err
	}
	return &model.GetBooksResponse{
		BookList: bookList,
		Total:    total,
		MaxPage:  total / size,
	}, nil
}

// GetInfiniteBooks get books service
func (s *BooksService) GetInfiniteBooks(ctx context.Context, req *model.GetInfiniteBooksRequest) (*model.GetInfiniteBooksResponse, error) {
	bookList, nextScrollId, err := s.repo.GetInfiniteBooks(ctx, req.ScrollId, req.Size)
	if err != nil {
		log.Printf("failed get infinite books: %v\n", err)
		return nil, err
	}
	return &model.GetInfiniteBooksResponse{
		NextScrollId: nextScrollId,
		BookList:     bookList,
	}, nil
}
