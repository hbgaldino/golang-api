package repository

import (
	"hbgaldino/golang-api/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	AllBook() []entity.Book
	FindBookById(id uint64) entity.Book
	InsertBook(b entity.Book) entity.Book
	DeleteBook(id uint64)
	UpdateBook(id uint64, b entity.Book) (entity.Book, error)
}

type bookRepository struct {
	connection *gorm.DB
}

func NewBookRepository(conn *gorm.DB) BookRepository {
	return &bookRepository{
		connection: conn,
	}
}

func (db *bookRepository) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Find(&books)
	return books
}

func (db *bookRepository) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	return b
}

func (db *bookRepository) FindBookById(id uint64) entity.Book {
	var book entity.Book
	db.connection.First(&book, id)
	return book
}

func (db *bookRepository) DeleteBook(id uint64) {
	db.connection.Delete(&entity.Book{}, id)
}

func (db *bookRepository) UpdateBook(id uint64, b entity.Book) (entity.Book, error) {
	var book entity.Book
	db.connection.First(&book, id)
	book.Title = b.Title
	book.Description = b.Description
	db.connection.Save(&book)
	return book, nil
}
