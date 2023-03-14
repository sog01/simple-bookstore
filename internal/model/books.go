package model

// Book represent a book model
type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	ISBN string `json:"isbn"`
}

// BookList represent a collection of book
type BookList []*Book

// PaginationByNumber represent a pagination by number for the collection of books
type PaginationByNumber struct {
	Total    int      `json:"total"`
	MaxPage  int      `json:"maxPage"`
	BookList BookList `json:"bookList"`
}

// PaginationByNumberArgs represent a pagination by number arguments
type PaginationByNumberArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// PaginationByCursor represent a pagination by cursor for the collection of books
type PaginationByCursor struct {
	NextCursor string   `json:"nextCursor"`
	BookList   BookList `json:"bookList"`
}

// PaginationByCursorArgs represent a pagination by cursor arguments
type PaginationByCursorArgs struct {
	NextCursor int `json:"nextCursor"`
	Size       int `json:"size"`
}
