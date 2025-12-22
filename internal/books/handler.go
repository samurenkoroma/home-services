package books

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"samurenkoroma/services/internal/app"
)

type BookHandler struct {
	router fiber.Router
}

func NewBookHandler(app *app.Polevod) {
	h := &BookHandler{
		router: app.App,
	}

	booksGroup := h.router.Group("/books")
	booksGroup.Get("/", h.getList)
	booksGroup.Get("/:id", h.getOne)
}

func (h BookHandler) getList(ctx *fiber.Ctx) error {

	var data = make([]Book, 15)

	for i := range data {
		data[i] = Book{
			Title:  fmt.Sprintf("Title %d", i),
			Author: "Author",
			Link:   "/awdawd",
		}

	}

	return ctx.JSON(data)
}

func (h BookHandler) getOne(ctx *fiber.Ctx) error {
	return ctx.JSON(
		Book{
			Title:  fmt.Sprintf("Title %d", ctx.Params("id")),
			Author: "Author",
			Link:   "/awdawd",
		},
	)
}
