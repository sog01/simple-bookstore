package repository

import "github.com/sog01/simple-bookstore/internal/model"

// SearchHits represent an object search result returned by Elasticsearch
type SearchHits struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Total Total  `json:"total"`
	Hits  []*Hit `json:"hits"`
}

type Total struct {
	Value int `json:"value"`
}

type Hit struct {
	Source *model.Book `json:"_source"`
}
