package books

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string
	Authors string
}
