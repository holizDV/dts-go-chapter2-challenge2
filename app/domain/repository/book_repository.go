package repository

import (
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/entity"
)

type BookRepository interface {
	CreateAt(b entity.Book) (entity.Book, error)
	Update(b entity.Book) (entity.Book, error)
	Delete(b entity.Book) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(id int) (entity.Book, error)
}

type repository struct {
}

func NewBookRepository() *repository {
	return &repository{}
}

func (r *repository) CreateAt(b entity.Book) (entity.Book, error) {
	return b, nil
}

func (r *repository) Update(b entity.Book) (entity.Book, error) {
	return b, nil
}

func (r *repository) Delete(b entity.Book) (entity.Book, error) {
	return b, nil
}

func (r *repository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	books = append(books, entity.Book{
		Id:          1,
		Title:       "Golang",
		Author:      "James",
		Description: "Google language",
	}, entity.Book{
		Id:          2,
		Title:       "Rust",
		Author:      "Cris Lawrence",
		Description: "Micro language",
	}, entity.Book{
		Id:          3,
		Title:       "Java",
		Author:      "Mark Buston",
		Description: "Learn of Java",
	})
	return books, nil
}

func (r *repository) FindByID(id int) (entity.Book, error) {
	book := entity.Book{
		Id:          id,
		Title:       "Golang",
		Author:      "James Cord",
		Description: "Build Clean Disc",
	}
	return book, nil
}
