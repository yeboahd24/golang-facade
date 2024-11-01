package repository

import (
	"github.com/yeboahd24/facade/pkg/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *model.Book) error
	GetAll() ([]model.Book, error)
	GetByID(id uint) (*model.Book, error)
	Update(book *model.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}
