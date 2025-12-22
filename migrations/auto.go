package main

import (
	"os"
	"samurenkoroma/services/internal/chitalka/books"
	"samurenkoroma/services/internal/link"
	"samurenkoroma/services/internal/stat"
	"samurenkoroma/services/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{}, &books.Book{})
}
