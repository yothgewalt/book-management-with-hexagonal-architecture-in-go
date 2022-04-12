package port

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null;unique;type:varchar(256)""`
	Author      string `gorm:"not null;type:varchar(64)""`
	Description string `gorm:"not null"`
}

type BookRepository interface {
	CreateBook(book Book) (*Book, error)
	GetAllBook() ([]*Book, error)
	GetBookById(id uint) (*Book, error)
	SoftDropBookById(id uint) error
	HardDropBookById(id uint) error
}
