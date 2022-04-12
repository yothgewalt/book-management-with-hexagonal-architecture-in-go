package service

import (
	"errors"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/port"
	"strings"
)

type bookServie struct {
	bookRepository port.BookRepository
}

func NewBookService(bookRepository port.BookRepository) BookService {
	return &bookServie{bookRepository: bookRepository}
}

func (b bookServie) NewBook(requester NewBookRequester) (*BookResponse, error) {
	requester.Title = strings.ToLower(requester.Title)

	if len(requester.Author) >= 5 && len(requester.Author) <= 24 {
		requester.Title = strings.ToLower(requester.Author)
	} else {
		return nil, errors.New("this author name must be more than 6 letters and must not be more than 24")
	}

	requester.Description = strings.ToLower(requester.Description)

	bookPayload := port.Book{
		Title:       requester.Title,
		Author:      requester.Author,
		Description: requester.Description,
	}
	book, err := b.bookRepository.CreateBook(bookPayload)
	if err != nil {
		return nil, err
	}

	responsePayload := BookResponse{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
	}

	return &responsePayload, nil
}

func (b bookServie) ReadBooks() ([]*BookResponse, error) {
	books, err := b.bookRepository.GetAllBook()
	if err != nil {
		return nil, err
	}

	var responses []*BookResponse
	for _, value := range books {
		responses = append(responses, &BookResponse{
			ID:          value.ID,
			CreatedAt:   value.CreatedAt,
			Title:       value.Title,
			Author:      value.Author,
			Description: value.Description,
		})
	}

	return responses, nil
}

func (b bookServie) ReadBookById(id uint) (*BookResponse, error) {
	book, err := b.bookRepository.GetBookById(id)
	if err != nil {
		return nil, err
	}

	response := BookResponse{
		ID:          book.ID,
		CreatedAt:   book.CreatedAt,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
	}

	return &response, nil
}

func (b bookServie) DeleteBookById(id uint) error {
	if err := b.bookRepository.DropBookById(id); err != nil {
		return err
	}

	return nil
}
