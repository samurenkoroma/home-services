package books

type File struct {
}

type BookResponse struct {
	Id      uint     `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type BookListResponse struct {
	Data  []BookResponse `json:"data"`
	Count int            `json:"count"`
}
