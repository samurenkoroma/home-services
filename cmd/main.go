package main

import (
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/app"
	"samurenkoroma/services/internal/auth"
	"samurenkoroma/services/internal/chitalka/books"
	"samurenkoroma/services/internal/user"
	"samurenkoroma/services/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	application := app.NewApplication(conf, database)

	//Репозитории
	userRepo := user.NewUserRepo(database)
	bookRepo := books.NewBookRepo(database)

	//Сервисы
	authService := auth.NewAuthService(userRepo)

	auth.NewAuthHandler(application.App, auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf.Auth,
	})

	books.NewBookHandler(books.BookHandlerDeps{Repo: bookRepo, Router: application.App})

	//home.NewHomeHandler(application)
	// pages.NewPageHandler(application)
	// weather.NewWeatherHandler(application)
	// finance.NewFinanceHandler(application)

	application.Run()
}
