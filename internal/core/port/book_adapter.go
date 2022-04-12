package port

import (
	"errors"
	"gorm.io/gorm"
)

type bookRepository struct {
	database *gorm.DB
}

func NewBookRepository(database *gorm.DB) BookRepository {
	return &bookRepository{database: database}
}

func (b bookRepository) CreateBook(book Book) (*Book, error) {
	if err := b.database.Table("books").Select("title").Where("title = ?", book.Title).Take(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx := b.database.Table("books").Model(&bookRepository{}).Create(&book).Error
			if tx != nil {
				return nil, err
			}

			return &book, nil
		} else {
			return nil, errors.New("this title could not be created because it is already in use")
		}
	}

	return nil, nil
}

func (b bookRepository) GetAllBook() ([]*Book, error) {
	var books []*Book
	if err := b.database.Table("books").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b bookRepository) GetBookById(id uint) (*Book, error) {
	var book *Book
	if err := b.database.Table("books").Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (b bookRepository) SoftDropBookById(id uint) error {
	var book *Book
	if err := b.database.Table("books").Select("id").Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b bookRepository) HardDropBookById(id uint) error {
	var book *Book
	if err := b.database.Table("books").Select("id").Unscoped().Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
