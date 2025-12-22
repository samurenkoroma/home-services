package books

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BookHandlerDeps struct {
	Repo   BookRepository
	Router fiber.Router
}

type BookHandler struct {
	repo   BookRepository
	router fiber.Router
}

func NewBookHandler(deps BookHandlerDeps) {
	h := &BookHandler{
		repo:   deps.Repo,
		router: deps.Router,
	}

	booksGroup := h.router.Group("/books")
	booksGroup.Get("", h.getList)
	booksGroup.Get("/:id", h.getOne)
}

func (h BookHandler) getList(ctx *fiber.Ctx) error {
	var params = new(BookQueryParams)
	if err := ctx.QueryParser(params); err != nil {
		return err
	}
	var books = h.repo.GetList(params)

	var data = []BookResponse{}

	for _, b := range books {
		data = append(data, BookResponse{
			Id:      b.ID,
			Title:   b.Title,
			Authors: strings.Split(b.Authors, ","),
		})
	}

	return ctx.JSON(BookListResponse{Data: data, Count: len(books)})
}

func (h BookHandler) getOne(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	book, err := h.repo.GetById(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(
		BookResponse{
			Title:   book.Title,
			Authors: strings.Split(book.Authors, ","),
			Id:      book.ID,
		},
	)
}
