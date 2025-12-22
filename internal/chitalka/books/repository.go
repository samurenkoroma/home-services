package books

import "samurenkoroma/services/pkg/db"

type BookRepository interface {
	Create(b *Book) (*Book, error)
	GetList(params *BookQueryParams) (books []Book)
	GetById(id uint) (*Book, error)
}

func NewBookRepo(database *db.Db) BookRepository {
	return &BookRepositoryImpl{
		database: database,
	}
}

type BookRepositoryImpl struct {
	database *db.Db
}

func (repo *BookRepositoryImpl) GetById(id uint) (book *Book, err error) {
	result := repo.database.DB.First(&book, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (repo *BookRepositoryImpl) Create(book *Book) (*Book, error) {
	result := repo.database.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

type BookQueryParams struct {
	Limit  int `query:"limit"`
	Cursor int `query:"cursor"`
}

func (repo *BookRepositoryImpl) GetList(params *BookQueryParams) (books []Book) {
	repo.database.
		Table("books").
		Where("deleted_at is null and id >= ?", params.Cursor).
		Order("id asc").
		Limit(params.Limit).
		Scan(&books)
	return books
}
