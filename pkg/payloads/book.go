package payloads

import (
	"fmt"
	"samurenkoroma/services/pkg/entities"
)

type BookRequest struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
type ResourceResponse struct {
	Type uint   `json:"type"`
	Link string `json:"link"`
}

type AuthorResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type BookResponse struct {
	Id        uint               `json:"id"`
	Title     string             `json:"title"`
	Authors   []AuthorResponse   `json:"authors"`
	Resources []ResourceResponse `json:"resources"`
}

type BookListResponse struct {
	Data  []BookResponse `json:"data"`
	Count int            `json:"count"`
}

func MakeBookResponse(book entities.Book) BookResponse {
	var resources = []ResourceResponse{}
	var authors = []AuthorResponse{}
	if len(book.Resources) != 0 {
		for _, r := range book.Resources {
			resources = append(resources, ResourceResponse{
				Type: uint(r.Type),
				Link: fmt.Sprintf("http://lab.note:8080/books/resource/%d", r.ID),
			})
		}
	}

	if len(book.Authors) != 0 {
		for _, a := range book.Authors {
			authors = append(authors, AuthorResponse{
				Id:   a.ID,
				Name: a.Name,
			})
		}
	}
	return BookResponse{
		Title:     book.Title,
		Authors:   authors,
		Resources: resources,
		Id:        book.ID,
	}
}
