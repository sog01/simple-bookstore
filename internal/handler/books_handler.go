package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sog01/simple-bookstore/internal/model"
	"github.com/sog01/simple-bookstore/internal/service"
)

// BooksHandler represent an rest handler object
type BooksHandler struct {
	booksService service.Books
}

// NewBooksHandler construct new rest object
func NewBooksHandler(bookService service.Books) *BooksHandler {
	return &BooksHandler{booksService: bookService}
}

// GetBooks is get books from rest handler
func (bh *BooksHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	req := model.GetBooksRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponseBadRequest(w, errors.New("invalid request"))
		return
	}

	bookList, err := bh.booksService.GetBooks(r.Context(), &req)
	if err != nil {
		writeResponseInternalError(w, err)
	}
	writeResponseOK(w, bookList)
}

// Router is rest handler router
func (bh *BooksHandler) Router(mux *http.ServeMux) {
	mux.HandleFunc("/books/list", bh.GetBooks)
}

// ListenAndServe is listening and serving
func (bh *BooksHandler) ListenAndServe() {
	mux := &http.ServeMux{}
	log.Println("listening server on port 8080")
	bh.Router(mux)
	http.ListenAndServe(":8080", mux)
}

func writeResponseOK(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, response)
}

func writeResponseBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	writeResponse(w, map[string]interface{}{
		"error": err,
	})
}

func writeResponseInternalError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	writeResponse(w, map[string]interface{}{
		"error": err,
	})
}

func writeResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
