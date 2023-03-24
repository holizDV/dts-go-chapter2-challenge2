package service

import (
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/entity"
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/repository"
)

type BookService interface {
	CreateAt(b entity.BookRequest) (entity.Book, error)
	Update(id int, b entity.BookRequest) (entity.Book, error)
	Delete(id int) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(id int) (entity.Book, error)
}

type service struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *service {
	return &service{repo}
}

func (s *service) CreateAt(b entity.BookRequest) (entity.Book, error) {
	book := entity.Book{
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}

	return book, nil
}

func (s *service) Update(id int, b entity.BookRequest) (entity.Book, error) {
	book := entity.Book{
		Id:          id,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}
	return book, nil

}

func (s *service) Delete(id int) (entity.Book, error) {
	book := entity.Book{
		Id: id,
	}
	return book, nil
}

func (s *service) FindAll() ([]entity.Book, error) {
	return s.repo.FindAll()
}

func (s *service) FindByID(id int) (entity.Book, error) {
	return s.repo.FindByID(id)
}
