package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sog01/simple-bookstore/internal/model"
	"github.com/sog01/simple-bookstore/internal/service"
)

// Rest represent an rest handler object
type Rest struct {
	booksService service.Books
}

// NewRest construct new rest object
func NewRest(bookService service.Books) *Rest {
	return &Rest{booksService: bookService}
}

// GetBooks is get books from rest handler
func (rest *Rest) GetBooks(w http.ResponseWriter, r *http.Request) {
	req := model.GetBooksRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponseBadRequest(w, errors.New("invalid request"))
		return
	}

	bookList, err := rest.booksService.GetBooks(r.Context(), &model.GetBooksRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		writeResponseInternalError(w, err)
	}
	writeResponseOK(w, bookList)
}

// Router is rest handler router
func (rest *Rest) Router(mux *http.ServeMux) {
	mux.HandleFunc("/books/list", rest.GetBooks)
}

// ListenAndServe is listening and serving
func (rest *Rest) ListenAndServe() {
	mux := &http.ServeMux{}
	log.Println("listening server on port 8080")
	rest.Router(mux)
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
