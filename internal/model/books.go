package model

// Book represent a book model
type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
}

// BookList represent a collection of book
type BookList []*Book

// GetBooksResponse represent a pagination by number for the collection of books
type GetBooksResponse struct {
	Total    int      `json:"total"`
	MaxPage  int      `json:"maxPage"`
	BookList BookList `json:"bookList"`
}

// GetBooksRequest represent a pagination by number arguments
type GetBooksRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// GetInfiniteBooksResponse represent a pagination by cursor for the collection of books
type GetInfiniteBooksResponse struct {
	NextCursor string   `json:"nextCursor"`
	BookList   BookList `json:"bookList"`
}

// GetInfiniteBooksArgs represent a pagination by cursor arguments
type GetInfiniteBooksArgs struct {
	NextCursor int `json:"nextCursor"`
	Size       int `json:"size"`
}
