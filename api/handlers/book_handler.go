package handlers

import (
	"errors"
	"log"
	"net/http"
	"samurenkoroma/services/pkg/entities"
	"samurenkoroma/services/pkg/payloads"
	"samurenkoroma/services/pkg/repositories"
	"samurenkoroma/services/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func Create(repo repositories.BookRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var payload payloads.BookRequest
		if err := ctx.BodyParser(&payload); err != nil {
			log.Println(err, ctx.Request())
			return response.ERROR(ctx, errors.New("wrong data"), http.StatusBadRequest)
		}

		var authors = []entities.Author{}
		for _, p := range payload.Authors {
			a := entities.Author{
				Name: p,
			}

			authors = append(authors, a)
		}

		book, err := repo.Create(&entities.Book{
			Title:   payload.Title,
			Authors: authors,
		})

		if err != nil {
			return response.ERROR(ctx, err, http.StatusConflict)
		}

		return response.JSON(ctx, payloads.MakeBookResponse(*book))

	}
}
func GetList(repo repositories.BookRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params = repositories.NewBookQueryParams()

		if err := ctx.QueryParser(params); err != nil {
			return response.ERROR(ctx, err, http.StatusBadRequest)
		}
		var books = repo.GetList(params)

		var data = []payloads.BookResponse{}
		for _, b := range books {
			data = append(data, payloads.MakeBookResponse(b))
		}

		return response.JSON(ctx, payloads.BookListResponse{Data: data, Count: len(books)})
	}
}

func GetOne(repo repositories.BookRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id")

		book, err := repo.GetById(uint(id))
		if err != nil {
			return response.ERROR(ctx, err, http.StatusNotFound)
		}

		return response.JSON(ctx, payloads.MakeBookResponse(*book))
	}
}
