package main

import (
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/app"
	"samurenkoroma/services/internal/auth"
	"samurenkoroma/services/internal/books"
	"samurenkoroma/services/internal/finance"
	"samurenkoroma/services/internal/pages"
	"samurenkoroma/services/internal/user"
	"samurenkoroma/services/internal/weather"
	"samurenkoroma/services/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	application := app.NewApplication(conf, database)

	//Репозитории
	userRepo := user.NewUserRepo(database)

	//Сервисы
	authService := auth.NewAuthService(userRepo)

	auth.NewAuthHandler(application.App, auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf.Auth,
	})

	//home.NewHomeHandler(application)
	books.NewBookHandler(application)
	pages.NewPageHandler(application)
	weather.NewWeatherHandler(application)
	finance.NewFinanceHandler(application)

	application.Run()
}
