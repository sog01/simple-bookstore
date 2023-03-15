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

// GetBooksRequest represent a pagination by number requests
type GetBooksRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// GetInfiniteBooksResponse represent a pagination by scroll for the collection of books
type GetInfiniteBooksResponse struct {
	NextScrollId string   `json:"nextScrollId"`
	BookList     BookList `json:"bookList"`
}

// GetInfiniteBooksRequest represent a pagination by scroll requests
type GetInfiniteBooksRequest struct {
	ScrollId string `json:"scrollId"`
	Size     int    `json:"size"`
}
