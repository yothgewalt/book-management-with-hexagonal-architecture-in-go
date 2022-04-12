package port

import (
	"time"
)

type Book struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	Title       string `gorm:"not null;unique;type:varchar(256)""`
	Author      string `gorm:"not null;type:varchar(64)""`
	Description string `gorm:"not null"`
}

type BookRepository interface {
	CreateBook(book Book) (*Book, error)
	GetAllBook() ([]*Book, error)
	GetBookById(id uint) (*Book, error)
	DropBookById(id uint) error
}
