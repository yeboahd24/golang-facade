package service

import (
	"fmt"
	"github.com/yeboahd24/facade/pkg/dto"
	"github.com/yeboahd24/facade/pkg/model"
	"github.com/yeboahd24/facade/pkg/repository"
)

type BookService interface {
	CreateBook(req *dto.CreateBookRequest) (*model.Book, error)
	GetAllBooks() ([]model.Book, error)
	GetBook(id uint) (*model.Book, error)
	UpdateBook(id uint, req *dto.UpdateBookRequest) (*model.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(req *dto.CreateBookRequest) (*model.Book, error) {
	book := &model.Book{
		Title:  req.Title,
		Author: req.Author,
		Price:  req.Price,
	}

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBook(id uint) (*model.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) UpdateBook(id uint, req *dto.UpdateBookRequest) (*model.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("book not found")
	}

	book.Title = req.Title
	book.Author = req.Author
	book.Price = req.Price

	if err := s.repo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
