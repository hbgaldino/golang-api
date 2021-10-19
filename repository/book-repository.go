package repository

import (
	"hbgaldino/golang-api/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	AllBook() []entity.Book
	FindBookById(id uint64) entity.Book
	InsertBook(b entity.Book) entity.Book
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
