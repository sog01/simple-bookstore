package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sog01/simple-bookstore/internal/model"
)

// Books represent books repository contract
type Books interface {
	GetBooks(ctx context.Context, from, size int) (model.BookList, int, error)
	GetInfiniteBooks(ctx context.Context, scrollId string, size int) (model.BookList, string, error)
}

// BooksRepository represent books repository object implementation
type BooksRepository struct {
	esBaseURL string
}

// NewBooksRepository construct new books repository object
func NewBooksRepository(esBaseURL string) *BooksRepository {
	return &BooksRepository{
		esBaseURL: esBaseURL,
	}
}

// GetBooks get data from repository
func (r *BooksRepository) GetBooks(ctx context.Context, from, size int) (model.BookList,
	int,
	error) {
	esReq := map[string]int{
		"size": 5,
	}
	if size > 0 {
		esReq["size"] = size
	}
	if from > 0 {
		esReq["from"] = from
	}

	url := fmt.Sprintf("%s/books/_search", r.esBaseURL)

	searchHits := SearchHits{}
	err := r.do(url, esReq, &searchHits)
	if err != nil {
		return nil, 0, fmt.Errorf("failed get books from repository: %v", err)
	}

	bookList := model.BookList{}
	for _, hit := range searchHits.Hits.Hits {
		bookList = append(bookList, hit.Source)
	}

	return bookList, searchHits.Hits.Total.Value, nil
}

// GetInfiniteBooks get infinite books from search after API
func (r *BooksRepository) GetInfiniteBooks(ctx context.Context, scrollId string, size int) (model.BookList,
	string,
	error) {
	esReq := map[string]interface{}{
		"size": 5,
	}
	if scrollId != "" {
		esReq["search_after"] = []interface{}{scrollId}
	}
	if size > 0 {
		esReq["size"] = size
	}
	esReq["sort"] = map[string]interface{}{
		"id": "asc",
	}

	url := fmt.Sprintf("%s/books/_search", r.esBaseURL)
	searchHits := SearchHits{}
	err := r.do(url, esReq, &searchHits)
	if err != nil {
		return nil, "", fmt.Errorf("failed get books from repository: %v", err)
	}

	bookList := model.BookList{}
	for _, hit := range searchHits.Hits.Hits {
		bookList = append(bookList, hit.Source)
	}

	lastHit := searchHits.Hits.Hits[len(searchHits.Hits.Hits)-1]
	nextScrollId := fmt.Sprintf("%v", lastHit.Sort[0])
	return bookList, nextScrollId, nil
}

func (r *BooksRepository) do(url string, dataReq interface{}, dst interface{}) error {
	byteReq, _ := json.Marshal(dataReq)

	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(byteReq))
	req.Header.Add("Content-type", "application/json")

	httpClient := http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
