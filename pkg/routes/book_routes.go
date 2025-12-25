package routes

import (
	"samurenkoroma/services/api/handlers"
	"samurenkoroma/services/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(router fiber.Router, repo repositories.BookRepository) {
	booksGroup := router.Group("/books")

	booksGroup.Get("", handlers.GetList(repo))
	booksGroup.Post("", handlers.Create(repo))
	booksGroup.Get("/resource/:id", handlers.GetResource(repo))
	booksGroup.Get("/:id", handlers.GetOne(repo))
}
