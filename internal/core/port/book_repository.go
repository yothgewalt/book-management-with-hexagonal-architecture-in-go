package port

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `gorm:"not null;unique;type:varchar(256)" json:"title" binding:"required"`
	Author string `gorm:"not null;type:varchar(64)" json:"author" binding:"required"`
}
