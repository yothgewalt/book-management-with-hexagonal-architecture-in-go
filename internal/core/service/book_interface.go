package service

import "time"

type NewBookRequester struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required,min=5,max=24"`
	Description string `json:"description" binding:"required"`
}

type BookResponse struct {
	ID          uint `json:"id"`
	CreatedAt   time.Time
	UpdateAt    time.Time
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type BookService interface {
	NewBook(requester NewBookRequester) (*BookResponse, error)
	ReadBooks() ([]*BookResponse, error)
	ReadBookById(id uint) (*BookResponse, error)
	SoftDeleteBookById(id uint) (*BookResponse, error)
	HardDeleteBookById(id uint) (*BookResponse, error)
}
